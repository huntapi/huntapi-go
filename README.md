# HuntAPI Go SDK

The official Go SDK for the [HuntAPI](https://www.huntapi.com) media download API.

Download video and audio from public media URLs, or extract their metadata, with one async REST API.

The SDK exposes **6 typed methods**, one per documented endpoint, plus a generic `Request()`. It has no dependencies outside the standard library.

## Requirements

- Go 1.22 or newer

## Installation

```sh
go get github.com/huntapi/huntapi-go
```

## Get your API key

[Sign up for HuntAPI](https://app.huntapi.com/register) or sign in, then copy your API key from the [dashboard](https://app.huntapi.com). `NewClient("", nil)` reads the `HUNTAPI_API_KEY` environment variable.

## Quick start

```go
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
	data, err := response.Map()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
```

Parameter structs mirror the [API reference](https://docs.huntapi.com). Required fields are plain values; optional fields are pointers — use `huntapi.String`, `huntapi.Float`, `huntapi.Int` and `huntapi.Bool`.

Decode into your own types with `response.Decode(&value)`.

## Async jobs

Download endpoints are asynchronous: they return a `job_id`, and jobs can run for up to 12 hours. `WaitForJob` polls `GET /v1/jobs/{id}` until a terminal status (`CompletedJob`, `Error`, `NotFound`, `BadRequest`):

```go
submitted, err := client.VideoDownload(ctx, &huntapi.VideoDownloadParams{Query: "https://www.youtube.com/watch?v=Z5YP8kznWZs", VideoQuality: huntapi.String("720p")})
if err != nil {
	log.Fatal(err)
}
var accepted struct{ JobID string `json:"job_id"` }
if err := submitted.Decode(&accepted); err != nil {
	log.Fatal(err)
}
job, err := client.WaitForJob(ctx, accepted.JobID, &huntapi.WaitForJobOptions{
	OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
})
```

## GET or POST

Methods use GET when the endpoint supports it (parameters in the query string); `WithMethod("POST")` sends them as a JSON body.

```go
response, err := client.AudioDownload(ctx, &huntapi.AudioDownloadParams{Query: "test"}, huntapi.WithMethod("POST"))
```

## Any endpoint

```go
response, err := client.Request(ctx, "GET", "/v1/usage", nil)
```

## Errors

```go
response, err := client.MediaMetadata(ctx, &huntapi.MediaMetadataParams{Query: "https://www.youtube.com/watch?v=Z5YP8kznWZs"})
if apiError, ok := huntapi.IsAPIError(err); ok {
	log.Printf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
}
```

| Error | When |
| --- | --- |
| `*ConfigurationError` | Missing API key, invalid options or unsupported method. |
| `*SerializationError` | A required parameter is missing or a value can't be sent. |
| `*TransportError` | Network failure, timeout or cancelled context (`errors.Is(err, context.DeadlineExceeded)`). |
| `*APIError` | Non-2xx status (`StatusCode`, `Message`, `Response`). |
| `*ResponseDecodingError` | A success response could not be decoded. |

## Retries

Requests are retried up to 3 times after a network error, a `408`, `500`, `502`, `503` or `504`, or a rate limit — which the API returns as `401` (or `429`) with a `Rate limit exceeded` message. Per-key quota errors (`…: quota daily`) are not retried, nor are `400`, `402`, `404` or client-side timeouts. The delay starts at 1 second and grows 1.5× per retry (capped at 30 seconds); a `Retry-After` header wins.

```go
client, err := huntapi.NewClient("", &huntapi.ClientOptions{
	MaxRetries: 5,                      // negative disables retries
	RetryDelay: 500 * time.Millisecond,
})

// or for one request
response, err := client.MediaMetadata(ctx, &huntapi.MediaMetadataParams{Query: "https://www.youtube.com/watch?v=Z5YP8kznWZs"}, huntapi.WithMaxRetries(0))
```

## Batch requests

`huntapi.Batch` runs a function per item with at most 10 calls in flight and returns one `BatchResult` per item, in input order. A failing item never stops the others.

```go
inputs := []string{"https://www.youtube.com/watch?v=Z5YP8kznWZs", "https://www.youtube.com/watch?v=jNQXAC9IVRw"}
results := huntapi.Batch(ctx, inputs, 5, func(ctx context.Context, value string) (*huntapi.Response, error) {
	return client.MediaMetadata(ctx, &huntapi.MediaMetadataParams{Query: value})
})
for _, result := range results {
	if result.Err != nil {
		log.Printf("%s: %v", inputs[result.Index], result.Err)
		continue
	}
	fmt.Println(string(result.Value.Body))
}
```

| Constant | Default | Description |
| --- | --- | --- |
| `DefaultConcurrency` | `10` | Tasks run at the same time by batch |
| `DefaultTimeout` | `120 * time.Second` | Request timeout |
| `MaxRetries` | `3` | Retry attempts after a network error, a 408/5xx or a rate limit |
| `RetryBackoffFactor` | `1.5` | Exponential backoff multiplier |
| `DefaultRetryDelay` | `time.Second` | Delay before the first retry |

## Timeouts and custom HTTP clients

```go
client, err := huntapi.NewClient("", &huntapi.ClientOptions{
	Timeout:    60 * time.Second, // default 120 s; negative disables
	HTTPClient: &http.Client{Transport: myTransport},
})
```

Every method also honours the deadline and cancellation of the `context.Context` you pass.

## Endpoints

6 methods, generated from the [OpenAPI document](https://app.huntapi.com/openapi).

| Method | HTTP | Description |
| --- | --- | --- |
| `AudioDownload` | GET / POST `/v1/audio/download` | Download audio from a URL via HuntAPI and receive a job id. |
| `Jobs` | GET `/v1/jobs/{id}` | Poll status, progress (0–100), and result of an async HuntAPI job. |
| `MediaMetadata` | GET `/v1/media/metadata` | Extract metadata from a media URL without downloading the file. |
| `Stats` | GET `/v1/stats` | Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 mea… |
| `Usage` | GET `/v1/usage` | Retrieve real-time usage statistics and credit balance for your Huntapi account. |
| `VideoDownload` | GET / POST `/v1/video/download` | Video Download |

## License

MIT — see [LICENSE](LICENSE).
