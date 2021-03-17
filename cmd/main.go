package main

import (
	"github.com/joho/godotenv"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/apiserver"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/repository"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := repository.Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Passw:        os.Getenv("DB_PASSW"),
		DatabaseName: os.Getenv("DB_NAME"),
	}

	repo, err := repository.New(dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	serverConfig := apiserver.Config{
		BindAddr: os.Getenv("BindAddr"),
	}

	s, err := apiserver.New(&serverConfig, repo)

	if err != nil {
		log.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
