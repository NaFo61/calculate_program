package main

import (
	"github.com/NaFo61/calculate_program/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
