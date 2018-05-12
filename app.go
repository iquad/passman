package main

import "go.uber.org/zap"

type App struct {
	Logger *zap.SugaredLogger
}

func createAppSingleton() (*App, error) {
	var err error
	err = nil

	app := &App{
		Logger: initLogger(),
	}

	return app, err
}

func (app *App) start() {
	app.Logger.Infow("app started")
}
