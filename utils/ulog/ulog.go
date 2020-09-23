package ulog

import "fmt"

func Debug(args ...interface{}) {
	defaultLogger.log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.log.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	defaultLogger.log.Debug(fmt.Sprintln(args...))
}

func Debugw(msg string, args ...interface{}) {
	defaultLogger.log.Debugw(msg, args...)
}

func Info(args ...interface{}) {
	defaultLogger.log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.log.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	defaultLogger.log.Info(fmt.Sprintln(args...))
}

func Infow(msg string, args ...interface{}) {
	defaultLogger.log.Infow(msg, args...)
}

func Warning(args ...interface{}) {
	defaultLogger.log.Warn(args...)
	ReportWarning(args...)
}

func Warningf(format string, args ...interface{}) {
	defaultLogger.log.Warnf(format, args...)
	ReportWarningf(format, args...)
}

func Warningln(args ...interface{}) {
	defaultLogger.log.Warn(fmt.Sprintln(args...))
	ReportWarningln(args...)
}

func Warningw(msg string, args ...interface{}) {
	defaultLogger.log.Warnw(msg, args...)
	ReportWarningw(msg, args...)
}

func Error(args ...interface{}) {
	defaultLogger.log.Error(args...)
	ReportError(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.log.Errorf(format, args...)
	ReportErrorf(format, args...)
}

func Errorln(args ...interface{}) {
	defaultLogger.log.Error(fmt.Sprintln(args...))
	ReportErrorln(args...)
}

func Errorw(msg string, args ...interface{}) {
	defaultLogger.log.Errorw(msg, args...)
	ReportErrorw(msg, args...)
}

func Fatal(args ...interface{}) {
	ReportFatal(args...)
	defaultLogger.log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	ReportFatalf(format, args...)
	defaultLogger.log.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	ReportFatalln(args...)
	defaultLogger.log.Fatal(fmt.Sprintln(args...))
}

func Fatalw(msg string, args ...interface{}) {
	ReportFatalw(msg, args...)
	defaultLogger.log.Fatalw(msg, args...)
}

func Panic(args ...interface{}) {
	ReportPanic(args...)
	defaultLogger.log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	ReportPanicf(format, args...)
	defaultLogger.log.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	ReportPanicln(args...)
	defaultLogger.log.Panic(fmt.Sprintln(args...))
}

func Panicw(msg string, args ...interface{}) {
	ReportPanicw(msg, args...)
	defaultLogger.log.Panicw(msg, args...)
}

func V(level int) bool {
	return unmarshalLevel(defaultLogger.level) >= unmarshalLevel("info")
}
