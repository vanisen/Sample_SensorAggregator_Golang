# Sample_Sensor_Aggregator

## Pre-requisite

* Run rabbitMQ(5672) and mongoDB(27017) on default ports
* Set GOPATH accordingly

### First run sensor simulator main (sensor write data every 10  seconds)
* go run $GOPATH/SensorAggregator/sensorSimulator/main.go

### Second run the Coordinator
This is central coordinator which collects data from all the queues of rabbitMQ and push them to mongoDB

* go run $GOPATH/coordinator/main.go


### Third supportive api to list the data of sensors

* start api

* go run $GOPATH/api/main.go
* localhost:5000/sensor
* localhost:5000/sensor/<sensor name>

