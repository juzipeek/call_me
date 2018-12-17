package comm_middleware

import (
	"net/http"
	
	"github.com/goadesign/goa"
	
	"context"
)

// LogRequest creates a request logger comm_middleware.
// This comm_middleware is aware of the RequestID comm_middleware and if registered after it leverages the
// request ID for logging.
// If verbose is true then the middlware logs the request and response bodies.
func NewEntryMiddleware() goa.Middleware {
	
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			goa.LogInfo(ctx, "NewEntryMiddleware entry")
			
			// real process
			err := h(ctx, rw, req)
			resp := goa.ContextResponse(ctx)
			code := resp.ErrorCode
			goa.LogInfo(ctx, "NewEntryMiddleware end", "status", resp.Status, "error", code,
				"bytes", resp.Length, "ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx))
			
			return err
		}
	}
}
