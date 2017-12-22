package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("URL Shortener API", func() {
	Title("The URL Shortener API")
	Description("An API for a URL Shortener")
	Contact(func() {
		Name("Example")
		Email("test@example.com")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/api/")
	Origin("*", func() {
		Headers("Content-Type")
		Methods("GET", "POST", "PATCH", "DELETE", "PUT", "OPTION")
	})
	Consumes("application/json")
	Produces("application/json")
})
