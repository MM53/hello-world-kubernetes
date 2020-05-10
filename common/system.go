package common

import (
	"log"
	"os"
)

func GetData() map[string]interface{} {
	data := make(map[string]interface{})
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Could not get pod name")
		log.Println(err)
	}

	data["hostname"] = hostname
	return data
}
