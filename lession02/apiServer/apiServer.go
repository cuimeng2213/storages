package main

import (
	"fmt"
	"net/http"
	"os"
	"storages/lession02/apiServer/heartbeat"
	"storages/lession02/apiServer/locate"
	"storages/lession02/apiServer/objects"
)

func main() {
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/object/", objects.Handler)
	go heartbeat.ListenHeartbeat()
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		fmt.Println("start server fialed")
	}

}
