package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("Listen api in port :", config.Port)

	r := router.Router(db)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
