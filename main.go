package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println("Listen api in port :", config.Port)

	r := router.Router()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
