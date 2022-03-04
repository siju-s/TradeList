package apihelpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	Data    interface{} `json:"data"`
	Message string
	status  int
}

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
