package main

import (
	"fmt"
	"net/http"
	"os"
	"storages/lession03/apiServer/heartbeat"
	"storages/lession03/apiServer/locate"
	"storages/lession03/apiServer/objects"
	"storages/lession03/apiServer/versions"
)

func main() {
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/object/", objects.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	go heartbeat.ListenHeartbeat()
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		fmt.Println("start server fialed")
	}

}
