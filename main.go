package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	r "favart-api/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	timeOut, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil || timeOut == 0 {
		timeOut = 15
	}

	idleTimeOut, err := strconv.Atoi(os.Getenv("IDLETIMEOUT"))
	if err != nil || idleTimeOut == 0 {
		idleTimeOut = 120
	}

	router := r.AppRouter()
	address := ":" + port

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: time.Duration(timeOut) * time.Second,
		ReadTimeout:  time.Duration(timeOut) * time.Second,
		IdleTimeout:  time.Duration(idleTimeOut) * time.Second,
	}

	fmt.Println("Starting a server at localhost:" + port)
	log.Fatal(srv.ListenAndServe())
}
