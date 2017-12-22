package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Url = MediaType("application/vnd.url+json", func() {
	Description("A URL")
	Attributes(func() {
		Attribute("id", Integer, "URL ID", func() {
			Example(1)
		})
		Attribute("url", String, "External url", func() {
			Example("example.com")
		})
		Attribute("path", String, "URL path key", func() {
			Example("path")
		})
	})

	View("default", func() {
		Attribute("id")
		Attribute("url")
		Attribute("path")
	})
})

var Analytics = MediaType("application/vnd.analytics+json", func() {
	Description("Url analytics")
	Attributes(func() {
		Attribute("hits", Integer, func() {
			Example(1)
		})
	})

	View("default", func() {
		Attribute("hits")
	})
})
