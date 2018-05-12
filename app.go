package main

import "go.uber.org/zap"

type App struct {
	Logger *zap.SugaredLogger
}

func (app *App) InitApp() error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	app.Logger = sugar

	return nil
}
