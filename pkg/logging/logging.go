package logging

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func Init() {
	var log = logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}
}