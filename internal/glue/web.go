package glue

import (
	"go-tellme/internal/module/client"
	"go-tellme/platform/routers"
	"net/http"
)

type webHandler struct {
	handler client.Handler
}

func (w webHandler) Routers() []*routers.Router {
	return []*routers.Router{
		{
			Method:        http.MethodGet,
			Path:          "/",
			Handler:       w.handler.PageHome,
			MiddlewareOne: w.handler.EnsureIndex(),
		},
		{
			Method:        http.MethodGet,
			Path:          "/about",
			Handler:       w.handler.PageAbout,
			MiddlewareOne: w.handler.EnsureNotLoggedIn(),
		},
		{
			Method:        http.MethodGet,
			Path:          "/login",
			Handler:       w.handler.PageLogin,
			MiddlewareOne: w.handler.EnsureLoggedIn(),
		},
		{
			Method:        http.MethodPost,
			Path:          "/login",
			Handler:       w.handler.PageLoginPerform,
			MiddlewareOne: w.handler.EnsureLoggedIn(),
		},
		{
			Method:        http.MethodGet,
			Path:          "/signup",
			Handler:       w.handler.PageSignup,
			MiddlewareOne: w.handler.EnsureLoggedIn(),
		},
		{
			Method:        http.MethodPost,
			Path:          "/signup",
			Handler:       w.handler.PageSignupPerform,
			MiddlewareOne: w.handler.EnsureLoggedIn(),
		},
		{
			Method:        http.MethodGet,
			Path:          "/logout",
			Handler:       w.handler.PageLogout,
			MiddlewareOne: w.handler.EnsureNotLoggedIn(),
		},
	}
}

func WebInit(handler client.Handler) client.Route {
	return &webHandler{handler: handler}
}
