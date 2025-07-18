package plugin

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(err error, msg string)
}
