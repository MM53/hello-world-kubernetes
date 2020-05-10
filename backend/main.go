package main

import (
	"encoding/json"
	"fmt"
	"hello-world-kubernetes/common"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/info", getInfo)
	log.Println("Start serving on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getInfo(w http.ResponseWriter, _ *http.Request) {
	response, err := json.Marshal(common.GetData())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
