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
	submitted, err := client.VideoDownload(ctx, &huntapi.VideoDownloadParams{Query: "https://www.youtube.com/watch?v=Z5YP8kznWZs", VideoQuality: huntapi.String("720p")})
	if err != nil {
		log.Fatal(err)
	}
	var accepted struct {
		JobID string `json:"job_id"`
	}
	if err := submitted.Decode(&accepted); err != nil {
		log.Fatal(err)
	}
	job, err := client.WaitForJob(ctx, accepted.JobID, &huntapi.WaitForJobOptions{
		OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(job.Body))
}
