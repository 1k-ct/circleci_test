package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/ping", Hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
