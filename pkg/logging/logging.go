package logging

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func init() {
	var log = logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		FullTimestamp: true,
	}

	log_file, err := os.OpenFile("../../logs/log.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(log_file)
	log.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(log)
}