package rabbitApi

import (
	"net/http"
	"encoding/json"
)

type Queue struct {
	Name string `json:name`
	VHost string `json:vhost`
}


func GetActiveList() []Queue{

	manager := "http://127.0.0.1:15672/api/queues/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", manager, nil)
	req.SetBasicAuth("guest", "guest")
	resp, _ := client.Do(req)

	value := make([]Queue, 0)
	json.NewDecoder(resp.Body).Decode(&value)
	return value
}