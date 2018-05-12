package main

import (
	"fmt"
	"os"
	"runtime"
)

var app *App

func main() {
	// how many cpu/thread can we use?
	// answer: all of them minus 1
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	app, err := createAppSingleton()
	if err != nil {
		fmt.Printf("failed to initiate application. going shutdown.")
		os.Exit(1)
	}

	app.start()

	os.Exit(0)
}
