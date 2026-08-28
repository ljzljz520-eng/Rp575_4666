package inbound

import (
	"encoding/json"
	"supplierhub/internal/model"
)

func Decode(raw []byte) (model.InboundOrder, error) {
	var o model.InboundOrder
	e := json.Unmarshal(raw, &o)
	return o, e
}
func FilterByWarehouse(items []model.InboundOrder, w string) []model.InboundOrder {
	out := []model.InboundOrder{}
	for _, x := range items {
		if x.Warehouse == w {
			out = append(out, x)
		}
	}
	return out
}
