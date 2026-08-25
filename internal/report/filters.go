package report

import "supplierhub/internal/model"

func InboundByStatus(xs []model.InboundOrder, status string) []model.InboundOrder {
	out := []model.InboundOrder{}
	for _, x := range xs {
		if x.Status == status {
			out = append(out, x)
		}
	}
	return out
}
func InboundByLines(xs []model.InboundOrder, min, max int) []model.InboundOrder {
	out := []model.InboundOrder{}
	for _, x := range xs {
		if x.Lines >= min && x.Lines <= max {
			out = append(out, x)
		}
	}
	return out
}
func QualityByGrade(xs []model.QualityResult, grade string) []model.QualityResult {
	out := []model.QualityResult{}
	for _, x := range xs {
		if x.Grade == grade {
			out = append(out, x)
		}
	}
	return out
}
func FailedQuality(xs []model.QualityResult) []model.QualityResult {
	out := []model.QualityResult{}
	for _, x := range xs {
		if !x.Passed {
			out = append(out, x)
		}
	}
	return out
}
func BillsByPeriod(xs []model.SettlementBill, p string) []model.SettlementBill {
	out := []model.SettlementBill{}
	for _, x := range xs {
		if x.Period == p {
			out = append(out, x)
		}
	}
	return out
}
func BillsAbove(xs []model.SettlementBill, n int64) []model.SettlementBill {
	out := []model.SettlementBill{}
	for _, x := range xs {
		if x.Amount > n {
			out = append(out, x)
		}
	}
	return out
}
