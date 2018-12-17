package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("call_me", func() {
	Title("call_me")
	Version("0.1")
	Scheme("http")
	Host("zzz.localhost")
})

var _ = Resource("call_me", func() {
	BasePath("/meet")
	DefaultMedia(SayHi)

	Action("dial", func() {
		Scheme("http")
		Routing(GET("/:uid"))
		Params(func() {
			Param("uid", Integer)
		})

		Payload(CallMePayLoad)
		Response(OK)
	})
})

var CallMePayLoad = Type("CallMePayLoad", func() {
	Attribute("name", String, "your name")
	Attribute("age", Integer, "your age")
	Attribute("gender", String, "your gender, male or female")
	Attribute("phone_no", String, "your phone no.")
})

// application/vnd.call_me 仅作为标识名字
var SayHi = MediaType("application/vnd.call_me", func() {
	Description("just say hi")
	TypeName("SayHiMedia")
	ContentType("application/json")
	Attributes(func() {
		Attribute("my_name", String, "my name")
		Attribute("my_age", Integer, "my age")
		Attribute("my_phone", String, "my phone no,call me now.")
	})
	View("default", func() {
		Attribute("my_name")
		Attribute("my_age")
		Attribute("my_phone")
	})
})
