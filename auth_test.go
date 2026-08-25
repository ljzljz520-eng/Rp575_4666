package main

import (
	"os"
	"supplierhub/internal/audit"
	"supplierhub/internal/auth"
	"supplierhub/internal/model"
	"supplierhub/internal/store"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	p := "a.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	a := audit.New(s)
	x := auth.New(s, a)
	if e := x.AddSupplier("s1", "Acme", "pw"); e != nil {
		t.Fatal(e)
	}
	if e := x.Grant("s1", model.Permission{Inbound: true}); e != nil {
		t.Fatal(e)
	}
	if _, e := x.Login("s1", "pw"); e != nil {
		t.Fatal(e)
	}
}
