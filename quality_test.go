package main

import (
	"os"
	"supplierhub/internal/audit"
	"supplierhub/internal/model"
	"supplierhub/internal/quality"
	"supplierhub/internal/store"
	"testing"
)

func TestWorkflowThree(t *testing.T) {
	p := "q.db"
	defer os.Remove(p)
	s, _ := store.Open(p)
	defer s.Close()
	x := quality.New(s, audit.New(s))
	if e := x.Add(model.QualityResult{ID: "q1", InboundID: "o1", SupplierID: "s1", Grade: "A", Passed: true}); e != nil {
		t.Fatal(e)
	}
	r, e := x.Get("q1", "s1")
	if e != nil || !r.Passed {
		t.Fatal(e)
	}
}
