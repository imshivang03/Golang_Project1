package main

import (
	"github.com/imshivang03/Golang_Project1/app"
	"github.com/imshivang03/Golang_Project1/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
