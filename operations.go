// Code generated from the HuntAPI OpenAPI document. DO NOT EDIT.

package huntapi

import "context"

// AudioDownload — Audio Download. Download audio from a URL via HuntAPI and receive a job id. Credit cost: 1 credit per call.
//
// GET / POST /v1/audio/download
func (c *Client) AudioDownload(ctx context.Context, params *AudioDownloadParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["audio_download"], params.params(), options)
}

// Jobs — Jobs. Poll status, progress (0–100), and result of an async HuntAPI job.
//
// GET /v1/jobs/{id}
func (c *Client) Jobs(ctx context.Context, params *JobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jobs"], params.params(), options)
}

// MediaMetadata — Media Metadata. Extract metadata from a media URL without downloading the file. Credit cost: 1 credit per call.
//
// GET /v1/media/metadata
func (c *Client) MediaMetadata(ctx context.Context, params *MediaMetadataParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["media_metadata"], params.params(), options)
}

// Stats — Stats. Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 means contact support to activate it.
//
// GET /v1/stats
func (c *Client) Stats(ctx context.Context, params *StatsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stats"], params.params(), options)
}

// Usage — Usage. Retrieve real-time usage statistics and credit balance for your Huntapi account.
//
// GET /v1/usage
func (c *Client) Usage(ctx context.Context, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usage"], nil, options)
}

// VideoDownload — Video Download. Credit cost: 1 credit per call.
//
// GET / POST /v1/video/download
func (c *Client) VideoDownload(ctx context.Context, params *VideoDownloadParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["video_download"], params.params(), options)
}

// endpoints lists every documented operation by its snake_case name.
var endpoints = map[string]endpoint{
	"audio_download": {path: "/v1/audio/download", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"jobs":           {path: "/v1/jobs/{id}", methods: []string{"GET"}, json: true, required: []string{"id"}},
	"media_metadata": {path: "/v1/media/metadata", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"stats":          {path: "/v1/stats", methods: []string{"GET"}, json: true, required: []string{}},
	"usage":          {path: "/v1/usage", methods: []string{"GET"}, json: true, required: []string{}},
	"video_download": {path: "/v1/video/download", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
}
