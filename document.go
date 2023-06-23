/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import "github.com/jung-kurt/gofpdf"

type Document struct {
	pdf *gofpdf.Fpdf

	Options                   *Options         `json:"options,omitempty"`
	Header                    *HeaderFooter    `json:"header,omitempty"`
	Footer                    *HeaderFooter    `json:"footer,omitempty"`
	Type                      string           `json:"type,omitempty" validate:"required,oneof=INVOICE RECEIPT"`
	Number                    string           `json:"version,omitempty" validate:"max=32"`
	Description               string           `json:"description,omitempty" validate:"max=1024"`
	Notes                     string           `json:"notes,omitempty"`
	Company                   *Contact         `json:"company,omitempty" validate:"required"`
	Customer                  *Contact         `json:"customer,omitempty" validate:"required"`
	Items                     []*Item          `json:"items,omitempty"`
	Date                      string           `json:"date,omitempty"`
	ValidityDate              string           `json:"validity_date,omitempty"`
	PaymentTerm               string           `json:"payment_term,omitempty"`
	AfterCommission           *AfterCommission `json:"after_commission,omitempty"`
	ConsumptionTax            *ConsumptionTax  `json:"consumption_tax,omitempty"`
	WithholdingTax            *WithholdingTax  `json:"withholding_tax,omitempty"`
	PaymentFree               *PaymentFree     `json:"payment_free,omitempty"`
	PaidAmount                *PaidAmount      `json:"paid_amount,omitempty"`
	InvoiceRegistrationNumber *string          `json:"invoice_registration_number,omitempty"`
}
