//go:generate goagen bootstrap -d github.com/juzipeek/call_me/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"app"
	"log_wrapper"
	"os"
	
	"github.com/sirupsen/logrus"
	"github.com/goadesign/goa/logging/logrus"
	"comm_middleware"
)

func main() {
	// Create service
	service := goa.New("call_me")
	
	// Mount comm_middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(comm_middleware.NewEntryMiddleware())
	
	// create new logger
	logger := log_wrapper.NewWriterLogger(os.Stdout, logrus.InfoLevel)
	logger.Info("logrus test")
	service.WithLogger(goalogrus.New(logger))
	
	// Mount "call_me" controller
	c := NewCallMeController(service)
	app.MountCallMeController(service, c)
	
	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
