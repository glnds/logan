package main

import (
	"fmt"
	"os"

	"github.com/glnds/logan/internal/logan"
	"go.uber.org/zap"
)

var logger *zap.Logger

var version, build, commit, date string

func main() {
	// 	conf := masl.GetConfig()
	// 	if conf.Debug {
	// 		logger = masl.GetLogger("debug")
	// 	} else {
	// 		logger = masl.GetLogger("info")
	// 	}
	logger = logan.GetLogger("info")
	// logger.Info("------------------ w00t w00t logan for you!?  ------------------")

	// flags := parseFlags(conf)
	logger.Info("Parsed the commandline flags")

	// DoLogan()
	return
}

// DoMasl Allow other tools to integrate with Masl to assume an AWS role
func DoLogan() {
	fmt.Println("No  masl for you! You don't have permissions to any account!")
	os.Exit(0)
}
