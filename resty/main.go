package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type UserData struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type APIResponse struct {
	Data UserData `json:"data"`
}

func main() {
	client := resty.New()

	// Retry configuration
	client.
		SetRetryCount(3).                   // set retry times to 3
		SetRetryWaitTime(10 * time.Second). // wait 10 seconds between each attempt
		AddRetryCondition(func(r *resty.Response, err error) bool {
			// Retry on any error or non-200 status code
			return err != nil || r.StatusCode() != 200
		}). // Retry condition looks at dif. aspects of request to determine if retry should be performed
		AddRetryHook(func(resp *resty.Response, err error) {
			log.Print("Retrying request... \n")
		}) // Retry hook will be executed prior to each "retry"

	// Perform the GET request
	resp, err := client.R().
		SetHeader("x-api-key", "reqres-free-v1").
		SetResult(&APIResponse{}).
		Get("https://reqres.in/api/users/2")

	if err != nil {
		log.Fatalf("Request failed after retries: %v", err)
	}

	// Extract and print the response
	apiResponse := resp.Result().(*APIResponse)

	fmt.Printf("ID: %d\n", apiResponse.Data.ID)
	fmt.Printf("Email: %s\n", apiResponse.Data.Email)
	fmt.Printf("First Name: %s\n", apiResponse.Data.FirstName)
	fmt.Printf("Last Name: %s\n", apiResponse.Data.LastName)
	fmt.Printf("Avatar: %s\n", apiResponse.Data.Avatar)
}
