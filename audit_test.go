package main

import (
	"os"
	"supplierhub/internal/audit"
	"supplierhub/internal/store"
	"testing"
)

func TestAuditLog(t *testing.T) {
	p := "l.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	a := audit.New(s)
	a.Record("s1", "login", "session")
	xs, e := a.List()
	if e != nil || len(xs) != 1 {
		t.Fatal(e)
	}
}
