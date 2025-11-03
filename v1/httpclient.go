package plugin

import (
	"context"
	"net/http"
	"time"
)

// HTTPCLientSetters defines methods for configuring and performing HTTP client operations.
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

// HTTPClientMethods defines HTTP client methods for sending requests with various HTTP methods.
type HTTPClientMethods interface {
	// Post performs an HTTP POST request with context, URL, body, response decoding, and headers.
	Post(
		ctx context.Context,
		url string,
		body any,
		out any,
		headers map[string]string,
	) (*http.Response, error)

	// Put performs an HTTP PUT request with context, URL, body, response decoding, and headers.
	Put(
		ctx context.Context,
		url string,
		body any,
		out any,
		headers map[string]string,
	) (*http.Response, error)

	// Delete performs an HTTP DELETE request with context, URL, response decoding, and headers.
	Delete(
		ctx context.Context,
		url string,
		out any,
		headers map[string]string,
	) (*http.Response, error)

	// DoRequest performs a general HTTP request with a specified method, URL, body, response decoding, and headers.
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
	HTTPClientMethods
}
