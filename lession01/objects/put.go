package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	fName := os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]
	log.Printf("create file: %s \n", fName)
	f, e := os.Create(fName)
	defer f.Close()
	if e != nil {
		log.Println("put file error: ", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.Copy(f, r.Body)
}
