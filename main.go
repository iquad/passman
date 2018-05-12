package main

import (
	"fmt"
	"os"
)

var app App

func main() {
	app = App{}
	err := app.InitApp()
	if err != nil {
		fmt.Printf("failed to initiate application. going shutdown.")
		os.Exit(1)
	}
}
