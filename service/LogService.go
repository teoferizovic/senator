package service

import (
	"fmt"
	"github.com/teoferizovic/senator/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitiLogger() {

	log.SetFormatter(&log.JSONFormatter{})

	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", config.GetEnvData("LOG_PATH")+"log", "%Y-%m-%d.%H:%M:%S"),
		rotatelogs.WithLinkName(config.GetEnvData("LOG_PATH")+"current"),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)

	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}

	log.SetOutput(writer)

	return

}
