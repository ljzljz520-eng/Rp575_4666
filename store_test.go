package main

import (
	"os"
	"supplierhub/internal/model"
	"supplierhub/internal/store"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := "test.db"
	defer os.Remove(p)
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.Put("suppliers", "s1", model.Supplier{ID: "s1", Name: "Acme", Active: true}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	var x model.Supplier
	if e = s.Get("suppliers", "s1", &x); e != nil || x.Name != "Acme" {
		t.Fatal(e, x)
	}
}
