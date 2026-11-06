// Package huntapi is the official Go SDK for the HuntAPI API (https://www.huntapi.com).
//
// Download video and audio from public media URLs, or extract their metadata, with one async REST API.
//
// Create one client and reuse it:
//
//	client, err := huntapi.NewClient("", nil) // reads HUNTAPI_API_KEY
//	response, err := client.MediaMetadata(ctx, ...)
//
// Every documented endpoint has a typed method; Client.Request reaches any other path.
package huntapi
