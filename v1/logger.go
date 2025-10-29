package plugin

// Logger provides an interface for logging messages of different severity levels.
//
// It is used to log debug, info, warning and error messages.
//
// Methods can be implemented by types that conform to this interface,
// allowing them to be used as a single source of logging functionality.
type Logger interface {
	//
	// Debug
	//
	// Prints the given message to the logger for debugging purposes.
	//
	// @param msg string The message to be printed to the logger.
	Debug(msg string)

	// Info
	//
	// Logs an info message using the provided logger.
	//
	// @param msg string The message to be logged by the logger.
	Info(msg string)

	// Warn
	//
	// Logs a warning message using the provided logger.
	//
	// @param msg string The message to be logged by the logger.
	// @return error nil if successful, non-nil otherwise.
	//
	// Logs a fatal error message with the given message.
	//
	// @param err error The error to be logged by the logger.
	// @param msg string The message to be logged by the logger.
	Warn(msg string)

	// Error reports an error with a given message.
	//
	// @param err  error     the underlying error
	// @param msg  string    the error message to display
	Error(err error, msg string)
}
