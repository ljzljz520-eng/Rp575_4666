package api

import "net/http"

func (x *API) routes() map[string]http.Handler { return map[string]http.Handler{"/": x, "/health": x} }
