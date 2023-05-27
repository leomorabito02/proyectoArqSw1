package main

import (
	"proyectoArqSw1/app"
	"proyectoArqSw1/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
