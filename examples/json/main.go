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
	response, err := client.MediaMetadata(ctx, &huntapi.MediaMetadataParams{Query: "https://www.youtube.com/watch?v=Z5YP8kznWZs"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
