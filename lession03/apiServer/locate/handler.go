package locate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("Lcoate %s \n", strings.Split(r.URL.EscapedPath(), "/")[2])
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(info)
	w.Write(data)
}
