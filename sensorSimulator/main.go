package main

import ("SensorAggregator/sensor"
	"fmt")

func main(){
	//sensor simulator
	finished := make(chan bool)
	go sensor.WriteMessage("sandy",1,10,finished)
	go sensor.WriteMessage("rocky" ,1 ,100,finished)
	go sensor.WriteMessage("lunar" , 10 ,30,finished)
	<-finished // can check for all sensors,
	fmt.Println("Sensor finished")
}
