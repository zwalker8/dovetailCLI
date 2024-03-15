package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func DecodeResponse[D any, E any](res *http.Response, dst *D, errorStruct *E) (*D, *E) {
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err := json.NewDecoder(res.Body).Decode(errorStruct)
		if err != nil {
			log.Fatal(err)
		}

		PrettyPrint(errorStruct)

		// fmt.Printf("%+v", errorStruct)
		return nil, errorStruct
	}

	err := json.NewDecoder(res.Body).Decode(dst)
	if err != nil {
		log.Fatal(err)
	}

	PrettyPrint(dst)
	// fmt.Printf("%+v", dst)

	return dst, nil
}

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
