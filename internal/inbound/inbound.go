package inbound

import (
	"context"
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
func (x *Service) Create(id, supplier, warehouse string, lines int) error {
	if lines <= 0 {
		return errors.New("lines")
	}
	return x.s.Put("inbounds", id, model.InboundOrder{ID: id, SupplierID: supplier, Warehouse: warehouse, Status: "received", Lines: lines})
}
func (x *Service) Get(ctx context.Context, id, supplier string) (model.InboundOrder, error) {
	var o model.InboundOrder
	if e := x.s.Get("inbounds", id, &o); e != nil {
		return o, e
	}
	if o.SupplierID != supplier {
		return o, errors.New("forbidden")
	}
	select {
	case <-ctx.Done():
		return o, ctx.Err()
	default:
		return o, nil
	}
}
func (x *Service) List(supplier string) ([]model.InboundOrder, error) {
	raw, e := x.s.List("inbounds")
	if e != nil {
		return nil, e
	}
	out := []model.InboundOrder{}
	for _, r := range raw {
		var o model.InboundOrder
		if json.Unmarshal(r, &o) == nil && o.SupplierID == supplier {
			out = append(out, o)
		}
	}
	return out, nil
}
