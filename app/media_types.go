// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "URL Shortener API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/rymccue/goa-url-shortener-api/design
// --out=$(GOPATH)/src/github.com/rymccue/goa-url-shortener-api
// --version=v1.2.0-dirty

package app

// Url analytics (default view)
//
// Identifier: application/vnd.analytics+json; view=default
type Analytics struct {
	Hits *int `form:"hits,omitempty" json:"hits,omitempty" xml:"hits,omitempty"`
}

// A URL (default view)
//
// Identifier: application/vnd.url+json; view=default
type URL struct {
	// URL ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// URL path key
	Path *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// External url
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}