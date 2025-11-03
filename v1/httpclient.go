package plugin

import (
	"context"
	"net/http"
	"time"
)

type HTTPCLientSetters interface {
	// SetBaseURL sets the base URL for the HTTP client.
	SetBaseURL(base string) HTTPClient
	// SetHeader sets a header for the HTTP client.
	SetHeader(key, value string) HTTPClient
	// SetTimeout sets the timeout for the HTTP client.
	SetTimeout(d time.Duration) HTTPClient
	// SetMaxRetries sets the maximum number of retries for the HTTP client.
	SetMaxRetries(n int) HTTPClient
	// SetBackoff sets the backoff duration for the HTTP client.
	SetBackoff(d time.Duration) HTTPClient
	// SetLogger sets the logger for the HTTP client.
	SetLogger(l Logger) HTTPClient
	// Get sends an HTTP GET request to the specified URL with optional headers and unmarshals the response into out.
	Get(ctx context.Context, url string, out any, headers map[string]string) (*http.Response, error)
}

type HttpClientMethods interface {
	// Post sends an HTTP POST request to the specified URL with the given body,
	// headers, and decodes the response into out.
	Post(
		ctx context.Context,
		url string,
		body any,
		out any,
		headers map[string]string,
	) (*http.Response, error)
	Put(
		ctx context.Context,
		url string,
		body any,
		out any,
		headers map[string]string,
	) (*http.Response, error)
	Delete(
		ctx context.Context,
		url string,
		out any,
		headers map[string]string,
	) (*http.Response, error)

	DoRequest(
		ctx context.Context,
		method, url string,
		body any,
		out any,
		headers map[string]string,
	) (*http.Response, error)
}

// HTTPClient is an interface for HTTP clients.
type HTTPClient interface {
	HTTPCLientSetters
	HttpClientMethods
}
