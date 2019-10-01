package main

import "github.com/hfantin/clientgo/app"

func main() {
	app := app.App{}
	app.Initialize("guest", "guest", "test")
	app.Run(":3000")
}
