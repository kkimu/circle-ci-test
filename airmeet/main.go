package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// Hello Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Get("/", Hello)
	e.Post("/events", RegisterEvent)
	e.Get("/events/:major", GetEventInfo)
	e.Delete("/events/:major", RemoveEvent)

	e.Post("/users/:major", RegisterUser)
	e.Get("/users/:major", GetParticipants)
	e.Delete("/users/:major/:id", RemoveUser)
	// Start server
	e.Run(standard.New(":3000"))
}
