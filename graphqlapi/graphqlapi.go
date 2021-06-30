package graphqlapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseStruct struct {
	Data struct {
		User struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"user"`
	} `json:"data"`
}

func QueryGraphQLEP(url string) ResponseStruct {

	fmt.Println("Querying graphql end point......")
	resp, err := http.Get(url)

	// request.Header.Add("content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseData ResponseStruct
	json.Unmarshal(body, &responseData)

	fmt.Println("Response data", responseData)
	return responseData
}
