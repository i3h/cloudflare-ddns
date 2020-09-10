package main

import (
	log "github.com/sirupsen/logrus"
)

func init_log() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
