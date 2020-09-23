package ulog

import (
	"fmt"
	"os"
	"path"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevels = []string{"debug", "info", "warn", "error", "panic", "fatal"}

func validLogLevel(level string) bool {
	for _, val := range logLevels {
		if val == level {
			return true
		}
	}
	return false
}

type pkgConfig struct {
	// Whether to log to console
	Console bool `mapstructure:"console"`

	// Directory of the log files. Default is the current working directory.
	Directory string `mapstructure:"directory"`

	// Name of this program. Used as the prefix of log filenames.
	Name string `mapstructure:"name"`

	// debug, warn, info, error, panic, fatal
	DefaultLevel string `mapstructure:"defaultLevel"`

	// Config of log file rotation.
	Rotation logRotation `mapstructure:"rotation"`
}

func (lc *pkgConfig) Validate() error {
	if _, err := os.Stat(lc.LogDirectory()); os.IsNotExist(err) {
		msg := fmt.Sprintf("Invalid LoggerConfig. Log Directory %s does not exist.", lc.Directory)
		return errors.Wrap(err, msg)
	}

	if lc.Name == "" {
		return errors.New("Invalid LoggerConfig. Log file name is not defined.")
	}

	// if DefaultLevel is empty, it is implied that DefaultLevel is `info`
	if lc.DefaultLevel != "" && !validLogLevel(lc.DefaultLevel) {
		return errors.New("Invalid LoggerConfig. Specified log `DefaultLevel` is not available.")
	}

	return nil
}

func (lc *pkgConfig) LogDirectory() string {
	if lc.Directory == "" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return dir
	}

	return lc.Directory
}

func (lc *pkgConfig) ErrorLogFile() string {
	errorFile := lc.Name + ".error.log"
	return path.Join(lc.LogDirectory(), errorFile)
}

func (lc *pkgConfig) LogFile() string {
	logFile := lc.Name + ".log"
	return path.Join(lc.LogDirectory(), logFile)
}

// Default values for log rotation
const (
	maxSizeMb  = 100
	maxBackups = 30
	maxAgeDays = 30
)

type logRotation struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated
	MaxSize int `mapstructure:"maxSize"`

	// MaxBackups is the maximum number of old log files to retain
	MaxBackups int `mapstructure:"maxBackups"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename
	MaxAge int `mapstructure:"maxAge"`

	// Compress rotated files with gzip
	Compress bool `mapstructure:"compress"`
}

var defaultLogRotation = logRotation{MaxSize: maxSizeMb, MaxBackups: maxBackups, MaxAge: maxAgeDays}

func newRotatedLogger(file string, rotation *logRotation) *lumberjack.Logger {
	if rotation == nil {
		rotation = &defaultLogRotation
	} else {
		// merge rotation config with default values.
		mergo.Merge(rotation, defaultLogRotation)
	}

	return &lumberjack.Logger{
		Filename:   file,
		MaxSize:    rotation.MaxSize,
		MaxBackups: rotation.MaxBackups,
		MaxAge:     rotation.MaxAge,
		Compress:   rotation.Compress,
	}
}
