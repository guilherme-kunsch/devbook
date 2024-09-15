package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnection = ""
	Port             = 0
)

func ToLoad() {
	var erro error

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("%s:%s/%s?charset=utf&parseTime=True&loc=local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
