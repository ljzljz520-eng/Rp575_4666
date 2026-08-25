package main

import (
	"context"
	"os"
	"supplierhub/internal/audit"
	"supplierhub/internal/inbound"
	"supplierhub/internal/quality"
	"supplierhub/internal/settlement"
	"supplierhub/internal/store"
	"supplierhub/internal/workflow"
	"testing"
	"time"
)

func TestBusinessChain49(t *testing.T) {
	p := "b49.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	a := audit.New(s)
	i := inbound.New(s, a)
	q := quality.New(s, a)
	st := settlement.New(s, a)
	i.Create("o49", "s49", "W", 1)
	c := workflow.New(i, q, st)
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()
	time.Sleep(time.Millisecond)
	_, e := c.Process(ctx, "o49", "s49")
	if e == nil {
		t.Fatalf("deadline was not propagated")
	}
}
