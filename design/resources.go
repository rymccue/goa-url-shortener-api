package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("shortener", func() {
	Action("create", func() {
		Routing(
			POST("/s"),
		)
		Description("Create a url")
		Payload(CreateLinkPayload)
		Response(Created)
		Response(InternalServerError)
	})

	Action("get", func() {
		Routing(
			GET("/s/:path"),
		)
		Description("Get a url")
		Params(func() {
			Param("path", String, "Path", func() {
				Example("path")
			})
		})
		Response(OK, Url)
		Response(InternalServerError)
		Response(NotFound)
	})

	Action("update", func() {
		Routing(
			PUT("/s/:path"),
		)
		Description("Update a url")
		Params(func() {
			Param("path", String, "Path", func() {
				Example("path")
			})
		})
		Payload(UpdateLinkPayload)
		Response(OK)
		Response(InternalServerError)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(
			DELETE("/s/:path"),
		)
		Description("Delete a url")
		Params(func() {
			Param("path", String, "Path", func() {
				Example("path")
			})
		})
		Response(OK)
		Response(InternalServerError)
		Response(NotFound)
	})

	Action("analytics", func() {
		Routing(
			GET("/s/:path/analytics"),
		)
		Description("Get analytics for a url")
		Params(func() {
			Param("path", String, "Path", func() {
				Example("path")
			})
		})
		Response(OK, Analytics)
		Response(InternalServerError)
		Response(NotFound)
	})
})
