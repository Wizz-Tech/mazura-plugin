package plugin

// Context represents an abstraction for handling HTTP requests and responses.
type Context interface {
	// JSON sends a JSON response with the specified status code and body.
	JSON(status int, body any) error

	// String sends a plain text response with the specified status code and message.
	String(status int, message string) error

	// Query retrieves the value of a query parameter by its key.
	Query(key string) string

	//
	Param(key string) string
}

// HandlerFunc defines a function type for handling HTTP requests using a context and returning an error, if any.
type HandlerFunc func(ctx Context) error

// Router defines methods for routing HTTP requests.
type Router interface {
	// Get registers a handler for HTTP GET requests to the specified``` pathgo.
	Get(path string, handler HandlerFunc)

	// Post registers a handler for HTTP POST requests to the specified path.
	Post(path string, handler HandlerFunc)

	// Put registers a handler for HTTP PUT requests to the specified path.
	Put(path string, handler HandlerFunc)

	// Delete registers a handler for HTTP DELETE requests to the specified path.
	Delete(path string, handler HandlerFunc)

	// Group creates a new router group with a specified prefix and returns the Router instance.
	Group(prefix string) Router
}
