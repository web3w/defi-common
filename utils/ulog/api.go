package ulog

// Logger interface is compatible with grpc-go grpclog.Loggerv2 and etcd raft.Logger interfaces
type Logger interface {
	// Debug logs are typically voluminous, and are usually disabled in production.
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Debugw(msg string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Infow(msg string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Warningw(msg string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Errorw(msg string, args ...interface{})

	// Panic logs a message, then panics.
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Panicw(msg string, args ...interface{})

	// Fatal logs a message, then calls os.Exit(1).
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Fatalw(msg string, args ...interface{})

	// V reports whether to output log messages at verbosity level l.
	V(l int) bool

	With(key string, value interface{}) Logger

	WithFields(fields map[string]interface{}) Logger
}

// NewLogger creates a logger with "info" level.
func NewLogger() Logger {
	return NewLoggerWithLevel(pkgCfg.DefaultLevel)
}

// NewLoggerWithLevel creates a ulog logger instance with custom log level. An "info" level logger
// will be created if `level` is "".
//
// Log levels available: "debug", "info", "warn", "error", "panic", "fatal".
func NewLoggerWithLevel(level string) Logger {
	return &ulogger{
		log: newProductionLogger(level),
		level: level,
	}
}
