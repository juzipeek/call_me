package main

import (
	"github.com/goadesign/goa"
	"app"
	"github.com/goadesign/goa/logging/logrus"
)

// CallMeController implements the call_me resource.
type CallMeController struct {
	*goa.Controller
}

// NewCallMeController creates a call_me controller.
func NewCallMeController(service *goa.Service) *CallMeController {
	return &CallMeController{Controller: service.NewController("CallMeController")}
}

// Dial runs the dial action.
func (c *CallMeController) Dial(ctx *app.DialCallMeContext) error {
	// CallMeController_Dial: start_implement
	
	// Put your logic here
	
	//entry := goalogrus.Entry(ctx)
	//entry.Infof("Dail action uid:%d", ctx.UID)
	
	goalogrus.Entry(ctx).Info("Dial action uid:", ctx.UID)
	
	res := &app.SayHiMedia{}
	age := 18
	name := "MayBe"
	phone := "13288882638"
	
	res.MyAge = &age
	res.MyName = &name
	res.MyPhone = &phone
	
	return ctx.OK(res)
	// CallMeController_Dial: end_implement
}
