package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huntapi/huntapi-go"
)

func main() {
	client, err := huntapi.NewClient("", nil) // reads HUNTAPI_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.Request(ctx, "GET", "/v1/usage", nil)
	if apiError, ok := huntapi.IsAPIError(err); ok {
		log.Fatalf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
