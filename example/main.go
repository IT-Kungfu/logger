package main

import (
	"github.com/IT-Kungfu/logger"
)

func main() {
	log, err := logger.New(&logger.Config{
		LogLevel:     "debug",
		SentryDSN:    "",
		LogstashAddr: "",
		ServiceName:  "logger",
		InstanceName: "dev",
	})
	if err != nil {
		panic(err)
	}

	log.Debugf("Debug message: %+v", log)
	log.Infof("Info message: %+v", log)
	log.Errorf("Error message: %+v", log)
}
