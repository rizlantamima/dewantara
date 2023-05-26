package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rizlantamima/dewantara/types"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUsers(w, r)
	}
}

func HandleUserDetail(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")

	id := r.URL.Path[0:3]
	fmt.Fprintf(w, "User ID: %s", id)
	// w.Write([]byte("Hello, World!"))

	// switch r.Method {
	// case http.MethodGet:
	// 	getUsers(w,r)
	// case http.MethodPost:
	// 	createUsers(w,r)
	// }
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	data := []types.User{}

	data = append(data, types.User{
		ID:       "1",
		FullName: "abang",
	})

	data = append(data, types.User{
		ID:       "2",
		FullName: "siapa aja boleh",
	})

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	data := types.User{
		ID:       "1",
		FullName: "abang",
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
