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
)

func main() {
	// 1. Connect to Database	
	db := driver.ConnectDB()

	// 2. Initialize Stores
	eStore := engineStore.New(db)
	cStore := carStore.New(db)

	// 3. Initialize Services
	eService := engineService.New(eStore)
	cService := carService.New(cStore)

	// 4. Initialize Handlers
	eHandler := engineHandler.New(eService)
	cHandler := carHandler.New(cService)

	// 5. Register Routes
	http.HandleFunc("POST /engines", eHandler.Create)
	http.HandleFunc("GET /engines/", eHandler.GetByID) // trailing slash for ID parsing
	
	http.HandleFunc("POST /cars", cHandler.Create)
	http.HandleFunc("GET /cars/", cHandler.GetByID)

	// 6. Start Server
	fmt.Println("Server starting on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}