package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {

	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(false)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	Logger.SetLevel(logrus.InfoLevel)
}
