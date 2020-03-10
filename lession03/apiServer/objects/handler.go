package objects

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	} else if m == http.MethodGet {
		get(w, r)
		return
	} else if m == http.MethodDelete {
		del(w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}