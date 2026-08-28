package main

import (
	"context"
	"os"
	"supplierhub/internal/audit"
	"supplierhub/internal/inbound"
	"supplierhub/internal/store"
	"testing"
)

func TestWorkflowTwo(t *testing.T) {
	p := "i.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	x := inbound.New(s, audit.New(s))
	x.Create("o1", "s1", "WH1", 3)
	o, e := x.Get(context.Background(), "o1", "s1")
	if e != nil || o.Lines != 3 {
		t.Fatal(e)
	}
}
