package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PrettyPrint(input any) {
	prettyFmt, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyFmt))
}

func GetAPIKEY() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return os.Getenv("API_KEY")
}
