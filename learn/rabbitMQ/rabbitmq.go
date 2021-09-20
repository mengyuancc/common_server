package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://mengyuan:garo469312@144.34.144.200:15672/mengyuan"
type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string
	Exchange string
	Key string
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbieMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq :=  &RabbitMQ{QueueName:queueName, Exchange:exchange, Key:key, Mqurl:MQURL}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接失败")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

//断开连接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//简单模式step:1.创建rabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return  NewRabbieMQ(queueName, "", "")
}

//简单模式step:2.生产消息
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//这样做是为了保证队列存在，消息能发送到队列中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
		)
	if err != nil {
		fmt.Println(err)
	}
	//2.发送消息到队列
	r.channel.Publish(
		r.Exchange,//如果是空的话 默认是default
		r.QueueName,
		//如果是true exchange无法找到队列会把消息返回给发送者
		false,
		//如果是true exchange检测队列没有消费者会把消息返回给发送者
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			Body:            []byte(message),
		},
		)
}




