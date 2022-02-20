package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"ventee-backend/configuration"
	"ventee-backend/handler"
)

func main() {
	env := os.Getenv("ENV")
	if env != configuration.ENVIRONMENTS.Heroku {
		dotenvError := godotenv.Load()
		if dotenvError != nil {
			log.Fatal(dotenvError)
		}
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = configuration.DEFAULT_PORT
	}

	http.HandleFunc("/", handler.Handle)

	log.Println("VENTEE-BACKEND is running on port", PORT)
	launchError := http.ListenAndServe(
		*flag.String(
			"ADDRESS",
			fmt.Sprintf(":"+PORT),
			"Bind an address",
		),
		nil,
	)
	if launchError != nil {
		log.Fatal(launchError)
	}
}
