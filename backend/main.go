package main

import (
	"repo/app"
	"repo/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
