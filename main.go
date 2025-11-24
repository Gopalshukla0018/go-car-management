package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gopalshukla0018/go-car-management/driver"
	carHandler "github.com/Gopalshukla0018/go-car-management/handler/car"
	engineHandler "github.com/Gopalshukla0018/go-car-management/handler/engine"
	carService "github.com/Gopalshukla0018/go-car-management/service/car"
	engineService "github.com/Gopalshukla0018/go-car-management/service/engine"
	carStore "github.com/Gopalshukla0018/go-car-management/store/car"
	engineStore "github.com/Gopalshukla0018/go-car-management/store/engine"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }


	db := driver.ConnectDB()


	eStore := engineStore.New(db)
	cStore := carStore.New(db)

	eService := engineService.New(eStore)
	cService := carService.New(cStore)

\
	eHandler := engineHandler.New(eService)
	cHandler := carHandler.New(cService)


	http.HandleFunc("POST /engines", eHandler.Create)
	http.HandleFunc("GET /engines/", eHandler.GetByID) 
	
	http.HandleFunc("POST /cars", cHandler.Create)
	http.HandleFunc("GET /cars/", cHandler.GetByID)

	
	fmt.Println("Server starting on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}