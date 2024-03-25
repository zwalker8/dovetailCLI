package api

import (
	"fmt"
	"net/http"
)

type TokenInfo struct {
	Data struct {
		ID        string `json:"id"`
		Subdomain string `json:"subdomain"`
	}
}

func (api *API) TokenInfo() (*TokenInfo, *APIError) {
	res := api.SendRequest(http.MethodGet, api.Routes.TokenInfo, nil)
	defer res.Body.Close()

	var apiResponse TokenInfo
	var apiError APIError

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (info *TokenInfo) Print() {
	fmt.Printf("ID: %v\n", info.Data.ID)
	fmt.Printf("Subdomain: %v\n\n", info.Data.Subdomain)
}
