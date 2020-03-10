package main

import (
	"net/http"
	"os"
	"storages/lession02/dataServer/heartbeat"
	"storages/lession02/dataServer/locate"
	"storages/lession02/dataServer/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		panic(err)
	}
}
