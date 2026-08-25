package api

import (
	"encoding/json"
	"net/http"
	"supplierhub/internal/audit"
	"supplierhub/internal/auth"
	"supplierhub/internal/inbound"
	"supplierhub/internal/quality"
	"supplierhub/internal/settlement"
)

type API struct {
	u *auth.Service
	i *inbound.Service
	q *quality.Service
	s *settlement.Service
	a *audit.Logger
}

func New(u *auth.Service, i *inbound.Service, q *quality.Service, s *settlement.Service, a *audit.Logger) *API {
	return &API{u: u, i: i, q: q, s: s, a: a}
}
func (x *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("content-type", "text/html")
		w.Write([]byte(page))
		return
	}
	if r.URL.Path == "/health" {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		return
	}
	http.NotFound(w, r)
}

const page = `<!doctype html><html><head><title>仓储供应商权限台</title></head><body><h1>仓储供应商权限台</h1><p>供应商入库、质检与结算权限管理</p></body></html>`
