package logging

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("example")

var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

/*
 * Below methods will be exported
 * Log.Debug, Log.Info, Log.Notice, Log.Warning, Log.Error, Log.Critical
 *
 */
func init() {
	// Init logging
	var logLevel logging.Level

	switch os.Getenv("LOG") {
	case "DEBUG":
		logLevel = logging.DEBUG
	default:
		logLevel = logging.INFO
	}

	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	logFormatter := logging.NewBackendFormatter(logBackend, logFormat)
	backendLeveled := logging.AddModuleLevel(logFormatter)
	backendLeveled.SetLevel(logLevel, "")
	logging.SetBackend(backendLeveled)
}

func Debug(i ...interface{}) {
	Log.Debug(i...)
}

func Info(i ...interface{}) {
	Log.Info(i...)
}

func Error(i ...interface{}) {
	Log.Error(i...)
}
