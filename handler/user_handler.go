package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rizlantamima/dewantara/types"
)


func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-type","application/json")
		getUsers(w,r)
	case http.MethodPost:
		createUsers(w,r)
	}	
}


func getUsers(w http.ResponseWriter, r *http.Request){
	data := []types.User{}

	data = append(data, types.User{
		ID: "1",
		FullName: "abang",
	})

	data = append(data, types.User{
		ID: "2",
		FullName: "siapa aja boleh",
	})

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}


func createUsers(w http.ResponseWriter, r *http.Request){
	data := types.User{
		ID: "1",
		FullName: "abang",
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}