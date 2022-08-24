//  https://github.com/wolfeidau/exitus
package main

// @title User API documentation
// @version 1.0.0
// @host localhost:5000
// @BasePath /api

import (
	"vmware.com/migrator-poc/db-sync-service/pkg/middleware"
)

func main() {
	// Initialize the database
	middleware.SetupRouter()
}
