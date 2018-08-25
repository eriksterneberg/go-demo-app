package logging

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("events-service")

var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Convention over configuration: drop uninteresting logs in the logging system
var logLevel = logging.DEBUG

func init() {
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	logFormatter := logging.NewBackendFormatter(logBackend, logFormat)
	backendLeveled := logging.AddModuleLevel(logFormatter)
	backendLeveled.SetLevel(logLevel, "")
	logging.SetBackend(backendLeveled)
}

var Debug = Log.Debug
var Debugf = Log.Debugf
var Info = Log.Info
var Infof = Log.Infof
var Warning = Log.Warning
var Warningf = Log.Warningf
var Error = Log.Error
var Errorf = Log.Errorf
