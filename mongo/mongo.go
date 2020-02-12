package mongo

import(
	"gopkg.in/mgo.v2"
	"fmt"
	"SensorAggregator/sensor"
)


func MongoConnect() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session, err
}

func GetSensorList() []string{
	fmt.Println("in sensor List")
	session,err := MongoConnect()
	list,err:=session.DB("rabbit").CollectionNames()
	fmt.Println(list,err)
	if err!=nil{
		fmt.Println(err)
	}
	session.Close()
	return list
}

func GetSensorData(name string) []sensor.Data{
	session,err := MongoConnect()
	var results []sensor.Data
	session.DB("rabbit").C(name).Find(nil).Limit(10).Sort("-$natural").All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}