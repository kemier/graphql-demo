// client.go
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graphql-demo/models"
	"io/ioutil"
	"net/http"
)

func Query() {
	client := &http.Client{}

	reqBody := []byte(`{
		"query": "query { user(id: \"1\") { id name email } }"
	}`)

	req, err := http.NewRequest("POST", "http://localhost:8080/graphql", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	var response struct {
		Data   map[string]models.User `json:"data"`
		Errors []Error                `json:"errors"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	if len(response.Errors) > 0 {
		fmt.Printf("Errors: %+v\n", response.Errors)
		return
	}

	fmt.Printf("User: %+v\n", response.Data["user"])
}

type Error struct {
	Message string `json:"message"`
}
