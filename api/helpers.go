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

func JoinQueryParams(page string, projects ...string) string {
	var pageParam string
	var projectParams string
	var projectQueryStrtings []string
	query := "?"
	seperator := "&"

	if page != "" {
		pageParam += fmt.Sprintf("page[start_cursor]=%v", page)
	}

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

	if page == "" || len(projects) == 0 {
		seperator = ""
		if page == "" && len(projects) == 0 {
			query = ""
		}
	}

	fmt.Println(query + pageParam + seperator + projectParams)

	return query + pageParam + seperator + projectParams
}
