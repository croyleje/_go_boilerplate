package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug.
	// Or this can be parsed from a config file as well via viper or similar
	// pacakges.
	if !ok {
		lvl = "debug"
	}
	// Parse string, this is built-in feature of logrus.
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// Set the global log level.
	logrus.SetLevel(ll)
}

func main() {
	logrus.Debug("Will only be visible if the loglevel permits it")
}
