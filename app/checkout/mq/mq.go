package mq

//
//var conn *amqp.Connection
//var ch *amqp.Channel
//var once sync.Once
//var respQueue amqp.Queue
//var err error
//
//func init() {
//	rand.Seed(time.Now().UTC().UnixNano())
//}
//func respQueueInit() {
//	once.Do(func() {
//		respQueue, err = ch.QueueDeclare(
//			"",    // 不指定队列名，默认使用随机生成的队列名
//			false, // durable
//			false, // delete when unused
//			true,  // exclusive
//			false, // noWait
//			nil,   // arguments
//		)
//		if err != nil {
//			panic(err)
//		}
//	})
//}
//func NewConnCh() {
//	var err error
//	conn, err = amqp.Dial("amqp://root:root@localhost:8080/")
//	if err != nil {
//		panic(err)
//	}
//	ch, err = conn.Channel()
//	if err != nil {
//		panic(err)
//	}
//}
//func ConnClose() {
//	ch.Close()
//	conn.Close()
//}
//func randomString(n int) string {
//	bytes := make([]byte, 1)
//	for i := 0; i < n; i++ {
//		bytes[i] = byte(randInt(65, 90))
//	}
//	return string(bytes)
//}
//func randInt(min int, max int) int {
//	return min + rand.Intn(max-min)
//}
//func PublishMessage(data []byte, exchangeName string, key string, mandatory bool, immediate bool) ([]byte, error) {
//	respQueueInit()
//	corrId := randomString(32)
//	err = ch.Publish(
//		exchangeName,
//		key,
//		mandatory,
//		immediate,
//		amqp.Publishing{
//			ContentType:   "text/plain",
//			ReplyTo:       respQueue.Name,
//			CorrelationId: corrId,
//			Body:          data,
//		},
//	)
//	if err != nil {
//		return nil, err
//	}
//	msgs, err := ch.Consume(
//		respQueue.Name, // queue
//		"",             // consumer
//		false,          // auto-ack
//		false,          // exclusive
//		false,          // no-local
//		false,          // no-wait
//		nil,            // args
//	)
//	if err != nil {
//		return nil, err
//	}
//	timer := time.NewTimer(3 * time.Second)
//	for {
//		select {
//		case d := <-msgs:
//			if d.CorrelationId == corrId {
//				// 响应id匹配，确认消息
//				err = d.Ack(false)
//				return d.Body, err
//			}
//		case <-timer.C:
//			return nil, errors.New("timeout")
//		}
//	}
//
//}
