package main

import app "GoPractice/GoCRUDs/internal/app"

func main() {
	a := app.App{}
	a.SetupService()
	a.StartHttpServer()
}
