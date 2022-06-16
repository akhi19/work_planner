package common

import (
	"os"

	"github.com/akhi19/work_planner/configs"
	log "github.com/sirupsen/logrus"
)

var logger log.Logger

func init() {
	f, err := os.OpenFile(configs.GetConfig().LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.SetLevel(log.WarnLevel)
}

func GetLogger() log.Logger {
	return logger
}
