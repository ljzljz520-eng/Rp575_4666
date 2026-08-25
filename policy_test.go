package main

import (
	"supplierhub/internal/auth"
	"supplierhub/internal/model"
	"testing"
)

func TestPermissionPolicy(t *testing.T) {
	p := model.Permission{Inbound: true}
	if !auth.CanRead(p, "inbound") || auth.CanRead(p, "quality") {
		t.Fatal("policy")
	}
}
