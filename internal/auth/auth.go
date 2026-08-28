package auth

import (
	"errors"
	"supplierhub/internal/audit"
	"supplierhub/internal/model"
	"supplierhub/internal/store"
)

type Service struct {
	s *store.Store
	a *audit.Logger
}

func New(s *store.Store, a *audit.Logger) *Service { return &Service{s: s, a: a} }
func (x *Service) AddSupplier(id, name, secret string) error {
	if id == "" || name == "" {
		return errors.New("invalid supplier")
	}
	return x.s.Put("suppliers", id, model.Supplier{ID: id, Name: name, Secret: secret, Active: true})
}
func (x *Service) Grant(id string, p model.Permission) error {
	var u model.Supplier
	if e := x.s.Get("suppliers", id, &u); e != nil {
		return e
	}
	p.SupplierID = id
	return x.s.Put("permissions", id, p)
}
func (x *Service) Login(id, secret string) (string, error) {
	var u model.Supplier
	if e := x.s.Get("suppliers", id, &u); e != nil {
		return "", e
	}
	if !u.Active || u.Secret != secret {
		return "", errors.New("unauthorized")
	}
	tok := "session-" + id
	if e := x.s.Put("sessions", tok, model.Session{Token: tok, SupplierID: id, Expires: 999999}); e != nil {
		return "", e
	}
	_ = x.a.Record(id, "login", "session")
	return tok, nil
}
func (x *Service) Resolve(tok string) (model.Session, error) {
	var z model.Session
	e := x.s.Get("sessions", tok, &z)
	return z, e
}
