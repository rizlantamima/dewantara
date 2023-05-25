package handler

import "net/http"

func HandleWellcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}