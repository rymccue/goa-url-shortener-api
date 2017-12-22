package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CreateLinkPayload = Type("CreateLinkPayload", func() {
	Attribute("path", String, func() {
		MinLength(6)
		MaxLength(100)
		Example("path")
	})

	Attribute("url", String, func() {
		MinLength(8)
		MaxLength(2000)
		Example("path")
	})
	Required("path", "url")
})

var UpdateLinkPayload = Type("UpdateLinkPayload", func() {
	Attribute("path", String, func() {
		MinLength(6)
		MaxLength(100)
		Example("path")
	})

	Attribute("url", String, func() {
		MinLength(8)
		MaxLength(2000)
		Example("path")
	})
	Required("path", "url")
})
