package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"storages/lession01/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
	fmt.Println("vim-go")
}
