package handlers

import (
	"net/http"
	"time"

	"github.com/pivotal-cf-experimental/lattice-app/routes"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/rata"
)

func New(logger lager.Logger) rata.Handlers {
	t := time.Now()
	handlers := rata.Handlers{
		routes.Env:   &Env{},
		routes.Hello: &Hello{Time: t},
		routes.Exit:  &Exit{Time: t},
	}

	for route, handler := range handlers {
		handlers[route] = &LoggingHandler{
			Route:   route,
			Handler: handler,
			Logger:  logger,
		}
	}

	return handlers
}

type LoggingHandler struct {
	Route   string
	Handler http.Handler
	Logger  lager.Logger
}

func (h *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.Logger.Session(h.Route)
	session.Info("request.begin")
	h.Handler.ServeHTTP(w, r)
	session.Info("request.end")
}
