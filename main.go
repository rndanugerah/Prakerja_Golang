package main

import (
	"prakerja/db"
	"prakerja/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":2703"))
}
