package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init_log() {
	if LOG_FILE != "" {
		file, _ := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
