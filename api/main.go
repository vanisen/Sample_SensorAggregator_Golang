package main

import (
	"net/http"
	"github.com/gorilla/mux"
	cont "SensorAggregator/api/controller"
)
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/sensor/{name}", cont.SensorData).Methods("GET")
	r.HandleFunc("/sensor/", cont.SensorList).Methods("GET")
	http.ListenAndServe(":5000", r)
}
