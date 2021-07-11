/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import (
	"fmt"
	"github.com/creasty/defaults"
	"github.com/jung-kurt/gofpdf"
)

type HeaderFooter struct {
	UseCustomFunc bool    `json:"-"`
	Text          string  `json:"text,omitempty"`
	FontSize      float64 `json:"font_size,omitempty" default:"7"`
	Pagination    bool    `json:"pagination,omitempty"`
}

type fnc func()

func (hf *HeaderFooter) ApplyFunc(pdf *gofpdf.Fpdf, fn fnc) {
	pdf.SetHeaderFunc(fn)
}

func (hf *HeaderFooter) applyHeader(d *Document, pdf *gofpdf.Fpdf) error {
	if err := defaults.Set(hf); err != nil {
		return err
	}

	if !hf.UseCustomFunc {
		pdf.SetHeaderFunc(func() {
			currentY := pdf.GetY()
			currentX := pdf.GetX()

			pdf.SetTopMargin(HeaderMarginTop)
			pdf.SetY(HeaderMarginTop)

			pdf.SetLeftMargin(BaseMargin)
			pdf.SetRightMargin(BaseMargin)

			pdf.SetFont("deja", "", hf.FontSize)
			_, lineHt := pdf.GetFontSize()
			html := pdf.HTMLBasicNew()
			html.Write(lineHt, hf.Text)
			if !hf.Pagination {
				pdf.AliasNbPages("")
				pdf.SetY(HeaderMarginTop + 8)
				pdf.SetX(195)
				pdf.CellFormat(10, 5, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()), "0", 0, "R", false, 0, "")
			}

			pdf.SetY(currentY)
			pdf.SetX(currentX)
			pdf.SetMargins(BaseMargin, BaseMarginTop, BaseMargin)
		})
	}

	return nil
}
