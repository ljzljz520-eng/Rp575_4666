package main

import (
	"context"
	"log"
	"net/http"
	"supplierhub/internal/api"
	"supplierhub/internal/audit"
	"supplierhub/internal/auth"
	"supplierhub/internal/inbound"
	"supplierhub/internal/quality"
	"supplierhub/internal/settlement"
	"supplierhub/internal/store"
)

func main() {
	s, err := store.Open("supplierhub.db")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	a := audit.New(s)
	users := auth.New(s, a)
	in := inbound.New(s, a)
	q := quality.New(s, a)
	st := settlement.New(s, a)
	h := api.New(users, in, q, st, a)
	srv := &http.Server{Addr: ":8080", Handler: h}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	_ = context.Background()
}
