// 定义包名为 model，表示这是一个用于模型定义的包
package model

// 导入所需的库
import (
	"context"                      // 提供上下文，用于取消操作或设置超时
	"encoding/json"                // JSON编码/解码库，用于将Go数据结构转换为JSON格式或反之
	"fmt"                          // 格式化I/O，用于格式化打印
	"github.com/redis/go-redis/v9" // Redis客户端库，用于与Redis交互
	"gorm.io/gorm"                 // GORM是一个流行的ORM库，用于数据库操作
	"time"                         // 时间库，用于处理时间相关的操作
)

// 定义产品结构体，包含基本信息以及关联的类别
type Product struct {
	Base                   // 继承自Base结构体，通常用于包含通用字段，如ID、CreatedAt等
	Name        string     `json:"name"`                                          // 产品名称，JSON标签用于序列化和反序列化
	Description string     `json:"description"`                                   // 产品描述
	Picture     string     `json:"picture"`                                       // 产品图片URL
	Price       float32    `json:"price"`                                         // 产品价格
	Categories  []Category `json:"categories" gorm:"many2many:product_category;"` // 产品所属的类别列表，many2many标签表示这是一个多对多的关系，中间表为product_category
}

// 定义产品查询结构体，包含上下文和数据库连接信息
type ProductQuery struct {
	ctx context.Context // 上下文，用于传递请求范围的数据、取消信号、超时信号和截止日期
	db  *gorm.DB        // 数据库连接对象，GORM的DB对象
}

// TableName 方法返回产品表的名称，这里返回 "product"
func (p Product) TableName() string {
	return "product"
}

// GetById 方法根据产品ID获取产品信息
func (p ProductQuery) GetById(productId int) (product Product, err error) {
	// 使用db对象从数据库中获取产品信息，WithContext方法用于传递上下文
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	// 返回产品信息和可能的错误
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) ProductQuery {
	return ProductQuery{ctx: ctx, db: db}
}

// SearchProducts 方法根据关键字搜索产品
func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	// 假设搜索条件是产品名称或描述包含关键字q
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	// 返回符合条件的产品列表和可能的错误
	return
}

// NewCachedProductQuery 方法创建并返回一个 CachedProductQuery 实例
func NewCachedProductQuery(pq ProductQuery, cacheClient *redis.Client) CachedProductQuery {
	// 初始化 CachedProductQuery 结构体，包含基础查询对象、Redis客户端和缓存键前缀
	return CachedProductQuery{productQuery: pq, cacheClient: cacheClient, prefix: "cloudwego_shop"}
}

// 定义 CachedProductQuery 结构体，包含用于查询和缓存操作的成员
type CachedProductQuery struct {
	productQuery ProductQuery  // 包含基础的 ProductQuery 对象
	cacheClient  *redis.Client // 包含 Redis 客户端对象
	prefix       string        // 缓存键的前缀，用于区分不同类型的缓存
}

// GetById 方法尝试从缓存中获取产品信息，如果缓存中没有则从数据库中查询并缓存结果
func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	// 构造缓存键，格式为 "prefix_product_by_id_productId"
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	// 尝试从Redis缓存中获取产品信息
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cacheKey)

	// 匿名函数用于处理缓存结果的解码
	err = func() error {
		// 如果获取缓存时发生错误，直接返回错误
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		// 将缓存结果转换为字节数组
		cachedResultByte, err2 := cachedResult.Bytes()
		if err2 != nil {
			return err2
		}
		// 将字节数组解码为Product结构体
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		// 如果解码成功，返回nil表示没有错误
		return nil
	}()
	// 如果上述步骤发生错误，则从数据库中查询产品信息
	if err != nil {
		// 调用基础的 GetById 方法从数据库中查询产品信息
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			// 如果查询数据库时发生错误，返回空的Product结构体和错误
			return Product{}, err
		}
		// 将查询到的产品信息编码为JSON格式
		encoded, err := json.Marshal(product)
		if err != nil {
			// 如果编码JSON时发生错误，返回查询到的产品信息和nil错误
			return product, nil
		}
		// 将编码后的JSON数据设置到Redis缓存中，有效期为1小时
		_ = c.cacheClient.Set(c.productQuery.ctx, cacheKey, encoded, time.Hour)
	}
	// 返回产品信息和可能的错误
	return
}

// GetProductById 方法根据产品ID从数据库中获取产品信息
func GetProductById(db *gorm.DB, ctx context.Context, productId int) (product Product, err error) {
	// 使用db对象从数据库中获取产品信息，WithContext方法用于传递上下文
	err = db.WithContext(ctx).Model(&Product{}).Where(&Product{Base: Base{ID: productId}}).First(&product).Error
	// 返回产品信息和可能的错误
	return product, err
}

// SearchProduct 方法根据关键字从数据库中搜索产品
func SearchProduct(db *gorm.DB, ctx context.Context, q string) (product []*Product, err error) {
	// 根据产品名称或描述包含关键字q进行查询
	err = db.WithContext(ctx).Model(&Product{}).Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	// 返回符合条件的产品列表和可能的错误
	return product, err
}
