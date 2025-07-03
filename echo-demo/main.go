package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance and counter
	var (
		e       = echo.New()
		counter = 0
	)

	// Middleware
	e.Use(middleware.Logger())

	e.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				counter++
				return next(c)
			}
		},
	)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			fmt.Sprintf("Number of requests: %v\n", counter),
		)
	})

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}
