package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeResponse[D any, E any](res *http.Response, dst *D, errorStruct *E) (*D, *E) {
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err := json.NewDecoder(res.Body).Decode(errorStruct)
		if err != nil {
			log.Fatal(err)
		}

		// PrettyPrint(errorStruct)

		// fmt.Printf("%+v", errorStruct)
		return nil, errorStruct
	}

	err := json.NewDecoder(res.Body).Decode(dst)
	if err != nil {
		log.Fatal(err)
	}

	// PrettyPrint(dst)
	// fmt.Printf("%+v", dst)

	return dst, nil
}
