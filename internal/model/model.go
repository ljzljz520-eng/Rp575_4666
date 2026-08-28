package model

type Supplier struct {
	ID, Name, Secret string
	Active           bool
}
type Permission struct {
	SupplierID                   string
	Inbound, Quality, Settlement bool
}
type InboundOrder struct {
	ID, SupplierID, Warehouse, Status string
	Lines                             int
}
type QualityResult struct {
	ID, InboundID, SupplierID, Grade string
	Passed                           bool
}
type SettlementBill struct {
	ID, SupplierID, Period string
	Amount                 int64
	Status                 string
}
type AuditLog struct{ ID, Actor, Action, Resource, At string }
type Session struct {
	Token, SupplierID string
	Expires           int64
}
