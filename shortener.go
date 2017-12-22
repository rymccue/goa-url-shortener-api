package main

import (
	"database/sql"
	"github.com/goadesign/goa"
	"github.com/rymccue/goa-url-shortener-api/app"
	"github.com/rymccue/goa-url-shortener-api/repositories"
)

// ShortenerController implements the shortener resource.
type ShortenerController struct {
	*goa.Controller
	DB *sql.DB
}

// NewShortenerController creates a shortener controller.
func NewShortenerController(service *goa.Service, db *sql.DB) *ShortenerController {
	return &ShortenerController{
		Controller: service.NewController("ShortenerController"),
		DB:         db,
	}
}

// Analytics runs the analytics action.
func (c *ShortenerController) Analytics(ctx *app.AnalyticsShortenerContext) error {
	// ShortenerController_Analytics: start_implement
	res, err := repositories.GetAnalytics(c.DB, ctx.Path)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		c.Service.LogError("Get Analytics", "err", err)
		return ctx.InternalServerError()
	}
	// ShortenerController_Analytics: end_implement
	return ctx.OK(res)
}

// Create runs the create action.
func (c *ShortenerController) Create(ctx *app.CreateShortenerContext) error {
	// ShortenerController_Create: start_implement
	p := ctx.Payload
	err := repositories.CreateURL(c.DB, p.URL, p.Path)
	if err != nil {
		c.Service.LogError("Create URL", "err", err)
		return ctx.InternalServerError()
	}
	// ShortenerController_Create: end_implement
	return ctx.Created()
}

// Delete runs the delete action.
func (c *ShortenerController) Delete(ctx *app.DeleteShortenerContext) error {
	// ShortenerController_Delete: start_implement
	err := repositories.DeleteURL(c.DB, ctx.Path)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		c.Service.LogError("Delete URL", "err", err)
		return ctx.InternalServerError()
	}
	// ShortenerController_Delete: end_implement
	return nil
}

// Get runs the get action.
func (c *ShortenerController) Get(ctx *app.GetShortenerContext) error {
	// ShortenerController_Get: start_implement
	res, err := repositories.GetURL(c.DB, ctx.Path)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		c.Service.LogError("Get URL", "err", err)
		return ctx.InternalServerError()
	}
	go repositories.AddUrlHit(c.DB, ctx.Path)
	// ShortenerController_Get: end_implement
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ShortenerController) Update(ctx *app.UpdateShortenerContext) error {
	// ShortenerController_Update: start_implement
	p := ctx.Payload
	err := repositories.UpdateURL(c.DB, p.URL, p.Path)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		c.Service.LogError("Update URL", "err", err)
		return ctx.InternalServerError()
	}
	// ShortenerController_Update: end_implement
	return nil
}
