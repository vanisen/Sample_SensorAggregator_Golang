package sensor

import (
	"github.com/streadway/amqp"
	"SensorAggregator/utils"
	"math/rand"
	"strconv"
	"time"
)
type Data struct{
	Value string `bson:"value"`
	At time.Time `bson:"time"`

}
func ConnectToRabbit() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	return conn,ch

}

func GetQueue ( ch *amqp.Channel ,name string) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name, //name
		false, //durable bool,
		false, //autoDelete bool,
		false, //exclusivebool,
		false, //noWait bool,
		nil)   //args amqp.Table)
	utils.FailOnError(err, "Failed to declare a queue")

	return &q
}

func WriteMessage(name string,min,max int, finished chan bool){
	conn, ch:= ConnectToRabbit()
	defer conn.Close()
	defer ch.Close()
	q:=GetQueue(ch,name)

	for {
		msg := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(giveValue(min,max)),
		}
		ch.Publish("", q.Name, false, false, msg)
		time.Sleep(10*time.Second)
	}
	finished <- true
}

func giveValue(min, max int) string{
	return strconv.Itoa(rand.Intn(max - min) + min)
}

