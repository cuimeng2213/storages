package main

import (
	"fmt"
)

func main() {
	http.HandleFunc("/locate/", locate.Handler)
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"))
	if err != nil {
		fmt.Println("start server fialed")
	}

}
