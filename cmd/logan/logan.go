package main

import (
	"fmt"
	"os"

	"github.com/glnds/masl/internal/masl"
	"go.uber.org/zap"
)

var logger *zap.Logger

var version, build, commit, date string

// Flags represents the command line flags
type Flags struct {
	Version     bool
	LegacyToken bool
	Profile     string
	Env         string
	Account     string
	Role        string
}

func main() {
	// 	conf := masl.GetConfig()
	// 	if conf.Debug {
	// 		logger = masl.GetLogger("debug")
	// 	} else {
	// 		logger = masl.GetLogger("info")
	// 	}

	logger.Info("------------------ w00t w00t logan for you!?  ------------------")

	// flags := parseFlags(conf)
	logger.Info("Parsed the commandline flags")

	DoLogan(conf, flags, password)
}

// DoMasl Allow other tools to integrate with Masl to assume an AWS role
func DoLogan(conf masl.Config, flags Flags, password string) {
	fmt.Println("No  masl for you! You don't have permissions to any account!")
	os.Exit(0)
}
