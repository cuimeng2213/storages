package objects

import (
	"lib/es"
	"lib/utils"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	//     object := strings.Split(r.URL.EscapedPath(), "/")[2]
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("Missing object hash in digist header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, e := storeObject(r.Body, url.PathEscape(hash))

	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size := utils.GetSizeFromHeader(r.Header)

	e = es.AddVersion(name, hash, size)

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
