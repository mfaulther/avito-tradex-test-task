package main

import (
	"github.com/joho/godotenv"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/apiserver"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := apiserver.Config{
		BindAddr:    os.Getenv("BindAddr"),
		DatabaseURL: os.Getenv("DatabaseURL"),
	}
	s, err := apiserver.New(&config)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
