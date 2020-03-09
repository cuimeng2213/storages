package objects

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func get() {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	readStream, err := getStream(object)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(w, readStream)
}
