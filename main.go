package main

import (
	"webapi/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := app.Init()
	app.Start()
}
