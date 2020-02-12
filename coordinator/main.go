package main

import (
	"SensorAggregator/rabbitApi"
	"log"
	"fmt"
	"SensorAggregator/sensor"
	"SensorAggregator/utils"
	"SensorAggregator/mongo"
	"time"
)


func ReadAllQueues(){
	q:= rabbitApi.GetActiveList()
	for _,v:=range q{
		go reader(v.Name)
	}

}

func main(){
	ReadAllQueues()
	var a string
	fmt.Scanln(&a)
}

func reader(name string) {
	conn, ch:= sensor.ConnectToRabbit()
	defer conn.Close()
	defer ch.Close()
	q:=sensor.GetQueue(ch,name)
	msgs, err := ch.Consume(
		q.Name, //queue string,
		"",     //consumer string,
		true,   //autoAck bool,
		false,  //exclusive bool,
		false,  //noLocal bool,
		false,  //noWait bool,
		nil)    //args amqp.Table)
	session,err:=mongo.MongoConnect()
	c := session.DB("rabbit").C(q.Name)
	c.DropCollection()


	utils.FailOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Received message with message: %s", msg.Body)
		c.Insert(sensor.Data{string(msg.Body), time.Now()})
	}
}