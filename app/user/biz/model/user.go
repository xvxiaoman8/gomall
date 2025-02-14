package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Create(ctx context.Context, db *gorm.DB, cache *redis.Client, user *User) error {
	// 先创建用户信息
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	// 再将用户信息添加到缓存中
	key := fmt.Sprintf("%s_%s_%s", "gomall", "get_user_by_email", user.Email)
	encoded, err := json.Marshal(user)
	if err != nil {
		return err
	}
	cache.Set(ctx, key, encoded, time.Hour)
	return nil
}

func Modify(ctx context.Context, db *gorm.DB, cache *redis.Client, user *User) error {
	// 先修改用户信息
	err := db.Model(&User{}).Where("email = ?", user.Email).Updates(user).Error
	if err != nil {
		return err
	}
	// 再将用户信息更新到缓存中
	key := fmt.Sprintf("%s_%s_%s", "gomall", "get_user_by_email", user.Email)
	encoded, err := json.Marshal(user)
	if err != nil {
		return err
	}
	cache.Set(ctx, key, encoded, time.Hour)
	return nil
}

func Delete(ctx context.Context, db *gorm.DB, cache *redis.Client, user *User) error {
	// 从数据库中删除用户
	err := db.Where("email = ?", user.Email).Delete(&User{}).Error
	if err != nil {
		return err
	}

	// 从缓存中删除用户信息
	key := fmt.Sprintf("%s_%s_%s", "gomall", "get_user_by_email", user.Email)
	err = cache.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetByEmail(ctx context.Context, db *gorm.DB, cache *redis.Client, email string) (*User, error) {
	// 先在缓存中查询
	var user User
	key := fmt.Sprintf("%s_%s_%s", "gomall", "get_user_by_email", email)
	result := cache.Get(ctx, key)
	err := func() error {
		err1 := result.Err()
		if err1 != nil {
			fmt.Println(err1)
			return err1
		}
		cachedResultByte, err2 := result.Bytes()
		if err2 != nil {
			fmt.Println(err2)
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &user)
		if err3 != nil {
			fmt.Println(err3)
			return err3
		}
		return nil
	}()
	if err != nil {
		// 查询不到再在数据库中查询
		err = db.Where("email = ?", email).First(&user).Error
		if err != nil {
			return nil, err
		}
		// 查询后将数据放入缓存
		key := fmt.Sprintf("%s_%s_%s", "gomall", "get_user_by_email", user.Email)
		encoded, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}
		cache.Set(ctx, key, encoded, time.Hour)
		return &user, err
	}
	return &user, err
}
