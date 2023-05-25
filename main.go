package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rizlantamima/dewantara/handler"
)

func main()  {
	listenAddressPort := flag.Int("port",3000,"port to serve HTTP on")
	flag.Parse()
	
	http.HandleFunc("/", handler.HandleWellcome)
	http.HandleFunc("/users", handler.HandleUsers)	


	listenAddressPortString := ":" + strconv.Itoa(*listenAddressPort)

	fmt.Printf("Your services are runing on port : %v",*listenAddressPort)
	http.ListenAndServe(listenAddressPortString,nil)
}