package controller

import (
	"net/http"

	"SensorAggregator/mongo"
	"github.com/gin-gonic/gin/json"
	"fmt"
	"github.com/gorilla/mux"
)

func SensorData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["name"]
	data := mongo.GetSensorData(param)
	bs,_:=json.Marshal(data)
	w.Write(bs)
}


func SensorList(w http.ResponseWriter, r *http.Request) {
	list:=mongo.GetSensorList()
	fmt.Println(list)
	bs,_:=json.Marshal(list)
	w.Write(bs)
}