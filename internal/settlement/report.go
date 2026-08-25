package settlement

import "supplierhub/internal/model"

func Total(items []model.SettlementBill) int64 {
	var n int64
	for _, x := range items {
		n += x.Amount
	}
	return n
}
func OpenBills(items []model.SettlementBill) []model.SettlementBill {
	out := []model.SettlementBill{}
	for _, x := range items {
		if x.Status == "open" {
			out = append(out, x)
		}
	}
	return out
}
