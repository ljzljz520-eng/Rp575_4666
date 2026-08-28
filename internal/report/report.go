package report

import (
	"fmt"
	"sort"
	"strings"
	"supplierhub/internal/model"
)

type Row struct {
	Kind, ID, Supplier, Status, Label string
	Amount                            int64
}
type Report struct {
	Rows   []Row
	Total  int64
	Counts map[string]int
}

func New() *Report { return &Report{Rows: []Row{}, Counts: map[string]int{}} }
func (r *Report) AddInbound(x model.InboundOrder) {
	r.Rows = append(r.Rows, Row{Kind: "inbound", ID: x.ID, Supplier: x.SupplierID, Status: x.Status, Label: x.Warehouse})
	r.Counts["inbound"]++
}
func (r *Report) AddQuality(x model.QualityResult) {
	label := x.Grade
	if x.Passed {
		label += " passed"
	} else {
		label += " review"
	}
	r.Rows = append(r.Rows, Row{Kind: "quality", ID: x.ID, Supplier: x.SupplierID, Status: label, Label: x.InboundID})
	r.Counts["quality"]++
}
func (r *Report) AddBill(x model.SettlementBill) {
	r.Rows = append(r.Rows, Row{Kind: "settlement", ID: x.ID, Supplier: x.SupplierID, Status: x.Status, Amount: x.Amount, Label: x.Period})
	r.Total += x.Amount
	r.Counts["settlement"]++
}
func (r *Report) Sort() {
	sort.Slice(r.Rows, func(i, j int) bool {
		if r.Rows[i].Kind == r.Rows[j].Kind {
			return r.Rows[i].ID < r.Rows[j].ID
		}
		return r.Rows[i].Kind < r.Rows[j].Kind
	})
}
func (r *Report) Filter(kind string) []Row {
	out := []Row{}
	for _, x := range r.Rows {
		if kind == "" || x.Kind == kind {
			out = append(out, x)
		}
	}
	return out
}
func (r *Report) CSV() string {
	r.Sort()
	b := strings.Builder{}
	b.WriteString("kind,id,supplier,status,label,amount\n")
	for _, x := range r.Rows {
		fmt.Fprintf(&b, "%s,%s,%s,%s,%s,%d\n", x.Kind, x.ID, x.Supplier, x.Status, x.Label, x.Amount)
	}
	return b.String()
}
func (r *Report) Summary() string {
	keys := []string{"inbound", "quality", "settlement"}
	parts := []string{}
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%d", k, r.Counts[k]))
	}
	return strings.Join(parts, " ") + fmt.Sprintf(" total=%d", r.Total)
}
func Build(in []model.InboundOrder, q []model.QualityResult, b []model.SettlementBill) *Report {
	r := New()
	for _, x := range in {
		r.AddInbound(x)
	}
	for _, x := range q {
		r.AddQuality(x)
	}
	for _, x := range b {
		r.AddBill(x)
	}
	return r
}
