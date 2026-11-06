// Code generated from the HuntAPI OpenAPI document. DO NOT EDIT.

package huntapi

// AudioDownloadParams holds the parameters of AudioDownload (GET / POST /v1/audio/download).
// Required fields are plain values; optional fields are pointers (use huntapi.String, huntapi.Float, huntapi.Bool) or nil slices and maps.
type AudioDownloadParams struct {
	// URL to download audio from.
	Query string `json:"query" api:"query,required"`
	// Desired download type. Supported values: audio (default). Default: "audio".
	DownloadType *string `json:"download_type" api:"download_type"`
	// Maximum duration of the downloaded content in seconds. If the audio is longer than this value, only the content up to this duration will be downloaded.
	MaxDuration *float64 `json:"max_duration" api:"max_duration"`
	// Desired audio quality. Default: "best".
	AudioQuality *string `json:"audio_quality" api:"audio_quality"`
	// Desired output format for the audio. Default: "mp3".
	AudioFormat *string `json:"audio_format" api:"audio_format"`
	// URL of the webhook where results will be sent. When specified, the system will attempt to send the response to this URL up to 3 times with a 5-second delay between each attempt. The results will be sent as a POST request.
	WebhookURL *string `json:"webhook_url" api:"webhook_url"`
	// Authorization header for the webhook. This parameter allows you to add authentication to your webhook endpoint.
	WebhookAuthorization *string `json:"webhook_authorization" api:"webhook_authorization"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *AudioDownloadParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 10)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.DownloadType != nil {
		values = append(values, param{"download_type", *p.DownloadType})
	}
	if p.MaxDuration != nil {
		values = append(values, param{"max_duration", *p.MaxDuration})
	}
	if p.AudioQuality != nil {
		values = append(values, param{"audio_quality", *p.AudioQuality})
	}
	if p.AudioFormat != nil {
		values = append(values, param{"audio_format", *p.AudioFormat})
	}
	if p.WebhookURL != nil {
		values = append(values, param{"webhook_url", *p.WebhookURL})
	}
	if p.WebhookAuthorization != nil {
		values = append(values, param{"webhook_authorization", *p.WebhookAuthorization})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// JobsParams holds the parameters of Jobs (GET /v1/jobs/{id}).
// Required fields are plain values; optional fields are pointers (use huntapi.String, huntapi.Float, huntapi.Bool) or nil slices and maps.
type JobsParams struct {
	// Job UUID returned as job_id by async endpoints such as /v1/video/download and /v1/audio/download.
	ID string `json:"id" api:"id,required"`
}

func (p *JobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 1)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	return values
}

// MediaMetadataParams holds the parameters of MediaMetadata (GET /v1/media/metadata).
// Required fields are plain values; optional fields are pointers (use huntapi.String, huntapi.Float, huntapi.Bool) or nil slices and maps.
type MediaMetadataParams struct {
	// Media URL to inspect (YouTube, Instagram, TikTok, Vimeo, X, Facebook, Twitch, Reddit, etc.).
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *MediaMetadataParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StatsParams holds the parameters of Stats (GET /v1/stats).
// Required fields are plain values; optional fields are pointers (use huntapi.String, huntapi.Float, huntapi.Bool) or nil slices and maps.
type StatsParams struct {
	// API family to check, for example video. Provide at least engine or path.
	Engine *string `json:"engine" api:"engine"`
	// Specific endpoint to check, for example /v1/video/download. If you also pass engine, use an endpoint from that API.
	Path *string `json:"path" api:"path"`
}

func (p *StatsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Engine != nil {
		values = append(values, param{"engine", *p.Engine})
	}
	if p.Path != nil {
		values = append(values, param{"path", *p.Path})
	}
	return values
}

// VideoDownloadParams holds the parameters of VideoDownload (GET / POST /v1/video/download).
// Required fields are plain values; optional fields are pointers (use huntapi.String, huntapi.Float, huntapi.Bool) or nil slices and maps.
type VideoDownloadParams struct {
	// Video URL to download.
	Query string `json:"query" api:"query,required"`
	// Desired download type. You have three options: Default: "audio_video".
	DownloadType *string `json:"download_type" api:"download_type"`
	// Maximum duration of the downloaded content in seconds. If the video/audio is longer than this value, only the content up to this duration will be downloaded.
	MaxDuration *float64 `json:"max_duration" api:"max_duration"`
	// Desired video quality. Available options: Default: "best".
	VideoQuality *string `json:"video_quality" api:"video_quality"`
	// Desired output format for the video. Available options: Default: "mp4".
	VideoFormat *string `json:"video_format" api:"video_format"`
	// URL of the webhook where results will be sent. When specified, the system will attempt to send the response to this URL up to 3 times with a 5-second delay between each attempt. The results will be sent as a POST request.
	WebhookURL *string `json:"webhook_url" api:"webhook_url"`
	// Authorization header for the webhook. This parameter allows you to add authentication to your webhook endpoint. This value will be included in the Authorization header when making the webhook request to your server.
	WebhookAuthorization *string `json:"webhook_authorization" api:"webhook_authorization"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *VideoDownloadParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 10)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.DownloadType != nil {
		values = append(values, param{"download_type", *p.DownloadType})
	}
	if p.MaxDuration != nil {
		values = append(values, param{"max_duration", *p.MaxDuration})
	}
	if p.VideoQuality != nil {
		values = append(values, param{"video_quality", *p.VideoQuality})
	}
	if p.VideoFormat != nil {
		values = append(values, param{"video_format", *p.VideoFormat})
	}
	if p.WebhookURL != nil {
		values = append(values, param{"webhook_url", *p.WebhookURL})
	}
	if p.WebhookAuthorization != nil {
		values = append(values, param{"webhook_authorization", *p.WebhookAuthorization})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}
