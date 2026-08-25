package quality

import (
	"encoding/json"
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
func (x *Service) Add(r model.QualityResult) error {
	if r.ID == "" || r.InboundID == "" {
		return errors.New("invalid result")
	}
	return x.s.Put("qualities", r.ID, r)
}
func (x *Service) Get(id, supplier string) (model.QualityResult, error) {
	var r model.QualityResult
	if e := x.s.Get("qualities", id, &r); e != nil {
		return r, e
	}
	if r.SupplierID != supplier {
		return r, errors.New("forbidden")
	}
	return r, nil
}
func (x *Service) List(supplier string) ([]model.QualityResult, error) {
	raw, e := x.s.List("qualities")
	if e != nil {
		return nil, e
	}
	out := []model.QualityResult{}
	for _, b := range raw {
		var r model.QualityResult
		if json.Unmarshal(b, &r) == nil && r.SupplierID == supplier {
			out = append(out, r)
		}
	}
	return out, nil
}
