package main

import (
	"github.com/scott-dn/go-boilerplate/api"
	"github.com/scott-dn/go-boilerplate/internal/app"
)

func main() {
	app := app.Init()
	api.StartHTTPServer(app)
}
