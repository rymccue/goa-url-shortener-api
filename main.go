//go:generate goagen bootstrap -d github.com/rymccue/goa-url-shortener-api/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/rymccue/goa-url-shortener-api/app"
	"github.com/rymccue/goa-url-shortener-api/utils/database"
)

func main() {
	// Create service
	service := goa.New("URL Shortener API")

	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), "url_shortener", os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		service.LogError("startup", "err", err)
	}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "shortener" controller
	c := NewShortenerController(service, db)
	app.MountShortenerController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
