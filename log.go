package main

import (
	"encoding/json"
	"io/ioutil"

	"runtime"

	"go.uber.org/zap"
)

// initiate logger or fallback
func initLogger() *zap.SugaredLogger {
	var sugar *zap.SugaredLogger

	// check config file exists
	f, err := ioutil.ReadFile("config/log.json")
	if err != nil {
		// failed to read file, use default config.
		return initDefaultLogger()
	}

	// config file imported, use it.
	var cfg zap.Config
	if err := json.Unmarshal(f, &cfg); err != nil {
		return initDefaultLogger()
	}

	if cfg.InitialFields == nil {
		cfg.InitialFields = make(map[string]interface{})
	}

	cfg.InitialFields["os"] = runtime.GOOS

	// build from given config
	logger, err := cfg.Build()

	if err != nil {
		return initDefaultLogger()
	}

	defer logger.Sync()
	sugar = logger.Sugar()

	return sugar
}

// fallback for default logger
func initDefaultLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger.Sugar()
}
