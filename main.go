package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rizlantamima/dewantara/handler"
)

func main()  {
	listenAddressPort := flag.Int("port",3000,"port to serve HTTP on")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleWellcome)
	mux.HandleFunc("/users", handler.HandleUsers)
	mux.HandleFunc("/users/", handler.HandleUserDetail)	


	listenAddressPortString := ":" + strconv.Itoa(*listenAddressPort)

	fmt.Printf("Your services are runing on port : %v",*listenAddressPort)
	log.Fatal(http.ListenAndServe(listenAddressPortString, mux))

}