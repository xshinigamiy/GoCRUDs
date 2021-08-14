package main

import (
	"GoCRUDs/internal/app"
)

func main() {
	a := app.App{}
	a.SetupService()
	a.StartHttpServer()
}
