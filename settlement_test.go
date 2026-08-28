package main

import (
	"os"
	"supplierhub/internal/audit"
	"supplierhub/internal/model"
	"supplierhub/internal/settlement"
	"supplierhub/internal/store"
	"testing"
)

func TestSettlementAccess(t *testing.T) {
	p := "s.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	x := settlement.New(s, audit.New(s))
	x.Add(model.SettlementBill{ID: "b1", SupplierID: "s1", Amount: 10})
	if _, e := x.Get("b1", "s2"); e == nil {
		t.Fatal("expected forbidden")
	}
}
