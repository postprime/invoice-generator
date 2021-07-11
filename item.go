/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import (
	"github.com/jung-kurt/gofpdf"
)

type Item struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Total int    `json:"total,omitempty"`
	Tax   *Tax   `json:"tax,omitempty"`
}

func (i *Item) appendColTo(options *Options, pdf *gofpdf.Fpdf) {
	baseY := pdf.GetY()

	pdf.SetX(ItemColNameOffset)
	pdf.MultiCell(ItemColUnitPriceOffset-ItemColNameOffset, 3, encodeString(i.Name), "", "", false)
	pdf.SetX(ItemColTotalHTOffset)
	pdf.CellFormat(ItemColTaxOffset-ItemColTotalHTOffset, 0, formatAmount(i.Total), "0", 0, "", false, 0, "")
	pdf.SetX(ItemColTotalTTCOffset)
	pdf.CellFormat(190-ItemColTotalTTCOffset, 0, formatAmount(i.Tax.Amount), "0", 0, "", false, 0, "")
	pdf.SetY(baseY + 0)
}
