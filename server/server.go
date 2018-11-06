package main

import (
	"net/http"
	"path/filepath"

	"github.com/labstack/echo"
)

// Text function to debug outputs
func test(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func main() {
	server := echo.New()

	// Set Static directory to server files from
	frontendRoot, err := filepath.Abs("../frontend")
	if err != nil {
		server.Logger.Fatal(err)
	}
	server.Static("/", frontendRoot)

	// Index page to point to handle server requests locally
	indexPath, err := filepath.Abs("../frontend/index.html")
	if err != nil {
		server.Logger.Fatal(err)
	}
	server.File("/", indexPath)

	// Routes to handle incoming requests
	server.GET("/v1/test", test)

	server.Logger.Fatal(server.Start(":8080"))
}
