package report

import (
	"fmt"
	"supplierhub/internal/model"
)

func FormatSupplier(x model.Supplier) string { return fmt.Sprintf("%s (%s)", x.Name, x.ID) }
func FormatPermission(x model.Permission) string {
	return fmt.Sprintf("inbound=%t quality=%t settlement=%t", x.Inbound, x.Quality, x.Settlement)
}
func FormatInbound(x model.InboundOrder) string {
	return fmt.Sprintf("%s %s %d", x.ID, x.Status, x.Lines)
}
func FormatQuality(x model.QualityResult) string {
	return fmt.Sprintf("%s %s %t", x.ID, x.Grade, x.Passed)
}
func FormatBill(x model.SettlementBill) string {
	return fmt.Sprintf("%s %d %s", x.ID, x.Amount, x.Status)
}
