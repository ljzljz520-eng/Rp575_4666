package main

import (
	"net/http/httptest"
	"os"
	"supplierhub/internal/api"
	"supplierhub/internal/audit"
	"supplierhub/internal/auth"
	"supplierhub/internal/inbound"
	"supplierhub/internal/quality"
	"supplierhub/internal/settlement"
	"supplierhub/internal/store"
	"testing"
)

func TestHTTPHealth(t *testing.T) {
	p := "h.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	a := audit.New(s)
	h := api.New(auth.New(s, a), inbound.New(s, a), quality.New(s, a), settlement.New(s, a), a)
	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
