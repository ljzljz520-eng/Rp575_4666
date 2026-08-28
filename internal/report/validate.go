package report

import (
	"errors"
	"supplierhub/internal/model"
)

func ValidateSupplier(x model.Supplier) error {
	if x.ID == "" {
		return errors.New("missing id")
	}
	if x.Name == "" {
		return errors.New("missing name")
	}
	if !x.Active {
		return errors.New("inactive")
	}
	return nil
}
func ValidatePermission(x model.Permission) error {
	if x.SupplierID == "" {
		return errors.New("missing supplier")
	}
	if !x.Any() {
		return errors.New("empty permission")
	}
	return nil
}
func ValidateInbound(x model.InboundOrder) error {
	if x.ID == "" || x.SupplierID == "" {
		return errors.New("identity")
	}
	if x.Lines < 1 {
		return errors.New("lines")
	}
	return nil
}
func ValidateQuality(x model.QualityResult) error {
	if x.ID == "" || x.InboundID == "" || x.SupplierID == "" {
		return errors.New("identity")
	}
	if x.Grade == "" {
		return errors.New("grade")
	}
	return nil
}
func ValidateBill(x model.SettlementBill) error {
	if x.ID == "" || x.SupplierID == "" {
		return errors.New("identity")
	}
	if x.Amount < 0 {
		return errors.New("amount")
	}
	return nil
}
