package router

import (
	"context"
	"net/http"
	"time"
)

type Logger interface {
	Printf(format string, v ...any)
}

type HTTPClient interface {
	SetBaseURL(base string) HTTPClient
	SetHeader(key, value string) HTTPClient
	SetTimeout(d time.Duration) HTTPClient
	SetMaxRetries(n int) HTTPClient
	SetBackoff(d time.Duration) HTTPClient
	SetLogger(l Logger) HTTPClient

	Get(ctx context.Context, url string, out any, headers map[string]string) (*http.Response, error)
	Post(ctx context.Context, url string, body any, out any, headers map[string]string) (*http.Response, error)
	Put(ctx context.Context, url string, body any, out any, headers map[string]string) (*http.Response, error)
	Delete(ctx context.Context, url string, out any, headers map[string]string) (*http.Response, error)
}
