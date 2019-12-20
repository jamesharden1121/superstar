package main

import (
	"superStar/bootstrap"
	"superStar/web/middleware/identity"
	"superStar/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("superStar", "6666666")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
