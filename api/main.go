package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/hello"
)

//HelloWorld return => "hello world, " + s(string)
func HelloWorld(s string) string {
	return "hello world, " + s
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", hello.Hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
