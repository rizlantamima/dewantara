package main

import (
	"flag"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main()  {
	listenAddressPort := flag.Int("port",3000,"port to serve HTTP on")
	flag.Parse()

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", handler.HandleWellcome)
	// mux.HandleFunc("/users", handler.HandleUsers)
	// mux.HandleFunc("/users/", handler.HandleUserDetail)	


	listenAddressPortString := ":" + strconv.Itoa(*listenAddressPort)

	app := echo.New()
	app.Logger.Fatal(app.Start(listenAddressPortString))

}