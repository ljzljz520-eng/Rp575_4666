package settlement

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
func (x *Service) Add(b model.SettlementBill) error {
	if b.Amount < 0 || b.ID == "" {
		return errors.New("invalid bill")
	}
	if b.Status == "" {
		b.Status = "open"
	}
	return x.s.Put("settlements", b.ID, b)
}
func (x *Service) Get(id, supplier string) (model.SettlementBill, error) {
	var b model.SettlementBill
	if e := x.s.Get("settlements", id, &b); e != nil {
		return b, e
	}
	if b.SupplierID != supplier {
		return b, errors.New("forbidden")
	}
	return b, nil
}
func (x *Service) List(supplier string) ([]model.SettlementBill, error) {
	raw, e := x.s.List("settlements")
	if e != nil {
		return nil, e
	}
	out := []model.SettlementBill{}
	for _, b := range raw {
		var x model.SettlementBill
		if json.Unmarshal(b, &x) == nil && x.SupplierID == supplier {
			out = append(out, x)
		}
	}
	return out, nil
}
