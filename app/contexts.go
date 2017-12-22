// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "URL Shortener API": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/rymccue/goa-url-shortener-api/design
// --out=$(GOPATH)/src/github.com/rymccue/goa-url-shortener-api
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// AnalyticsShortenerContext provides the shortener analytics action context.
type AnalyticsShortenerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Path string
}

// NewAnalyticsShortenerContext parses the incoming request URL and body, performs validations and creates the
// context used by the shortener controller analytics action.
func NewAnalyticsShortenerContext(ctx context.Context, r *http.Request, service *goa.Service) (*AnalyticsShortenerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AnalyticsShortenerContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPath := req.Params["path"]
	if len(paramPath) > 0 {
		rawPath := paramPath[0]
		rctx.Path = rawPath
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AnalyticsShortenerContext) OK(r *Analytics) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.analytics+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *AnalyticsShortenerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *AnalyticsShortenerContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CreateShortenerContext provides the shortener create action context.
type CreateShortenerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateLinkPayload
}

// NewCreateShortenerContext parses the incoming request URL and body, performs validations and creates the
// context used by the shortener controller create action.
func NewCreateShortenerContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateShortenerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateShortenerContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateShortenerContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateShortenerContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteShortenerContext provides the shortener delete action context.
type DeleteShortenerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Path string
}

// NewDeleteShortenerContext parses the incoming request URL and body, performs validations and creates the
// context used by the shortener controller delete action.
func NewDeleteShortenerContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteShortenerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteShortenerContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPath := req.Params["path"]
	if len(paramPath) > 0 {
		rawPath := paramPath[0]
		rctx.Path = rawPath
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteShortenerContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteShortenerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteShortenerContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// GetShortenerContext provides the shortener get action context.
type GetShortenerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Path string
}

// NewGetShortenerContext parses the incoming request URL and body, performs validations and creates the
// context used by the shortener controller get action.
func NewGetShortenerContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetShortenerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetShortenerContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPath := req.Params["path"]
	if len(paramPath) > 0 {
		rawPath := paramPath[0]
		rctx.Path = rawPath
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetShortenerContext) OK(r *URL) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.url+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetShortenerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetShortenerContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// UpdateShortenerContext provides the shortener update action context.
type UpdateShortenerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Path    string
	Payload *UpdateLinkPayload
}

// NewUpdateShortenerContext parses the incoming request URL and body, performs validations and creates the
// context used by the shortener controller update action.
func NewUpdateShortenerContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateShortenerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateShortenerContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPath := req.Params["path"]
	if len(paramPath) > 0 {
		rawPath := paramPath[0]
		rctx.Path = rawPath
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateShortenerContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateShortenerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateShortenerContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
