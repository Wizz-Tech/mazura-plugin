package plugin

type Context interface {
	JSON(status int, body any) error
	String(status int, message string) error
	Query(key string) string
	Param(key string) string
}

type HandlerFunc func(ctx Context) error

// Router defines methods for routing HTTP requests.
type Router interface {
	Get(path string, handler HandlerFunc)
	Post(path string, handler HandlerFunc)
	Put(path string, handler HandlerFunc)
	Delete(path string, handler HandlerFunc)
	Group(prefix string) Router
}
