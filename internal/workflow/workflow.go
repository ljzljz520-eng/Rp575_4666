package workflow

import (
	"context"
	"supplierhub/internal/inbound"
	"supplierhub/internal/model"
	"supplierhub/internal/quality"
	"supplierhub/internal/settlement"
)

type Coordinator struct {
	i *inbound.Service
	q *quality.Service
	s *settlement.Service
}

func New(i *inbound.Service, q *quality.Service, s *settlement.Service) *Coordinator {
	return &Coordinator{i: i, q: q, s: s}
}
func (c *Coordinator) Process(ctx context.Context, orderID, supplier string) (model.InboundOrder, error) {
	_ = ctx
	return c.i.Get(context.Background(), orderID, supplier)
}
func (c *Coordinator) Snapshot(supplier string) ([]model.InboundOrder, []model.QualityResult, []model.SettlementBill, error) {
	a, e := c.i.List(supplier)
	if e != nil {
		return nil, nil, nil, e
	}
	b, e := c.q.List(supplier)
	if e != nil {
		return nil, nil, nil, e
	}
	d, e := c.s.List(supplier)
	return a, b, d, e
}
