package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func DecodeResponse[D any, E any](res *http.Response, dst *D, errorStruct *E) (*D, *E) {
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err := json.NewDecoder(res.Body).Decode(errorStruct)
		if err != nil {
			log.Fatal(err)
		}

		return nil, errorStruct
	}

	err := json.NewDecoder(res.Body).Decode(dst)
	if err != nil {
		log.Fatal(err)
	}

	return dst, nil
}

func JoinQueryParams(page string, limit uint8, projects ...string) string {
	var limitParam string
	var pageParam string
	var projectParams string
	var projectQueryStrtings []string
	var querylist []string

	if page != "" {
		pageParam += fmt.Sprintf("page[start_cursor]=%v", page)
	}

	if limit > 100 {
		limit = 100
	}

	limitParam += fmt.Sprintf("page[limit]=%v", limit)

	if len(projects) != 0 {
		for index, id := range projects {
			if len(projects) == 1 {
				projectQueryStrtings = append(projectQueryStrtings, fmt.Sprintf("filter[project_id]=%v", id))
			} else {
				projectQueryStrtings = append(projectQueryStrtings, fmt.Sprintf("filter[project_id][%d]=%v", index, id))
			}
		}
		projectParams += strings.Join(projectQueryStrtings, "&")
	}

	querylist = append(querylist, pageParam, limitParam, projectParams)
	queryString := "?"
	for i, param := range querylist {
		if param != "" {
			queryString += param
			if i != len(querylist)-1 {
				queryString += "&"
			}
		}
	}

	return queryString
}
