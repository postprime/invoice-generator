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

func (hf *HeaderFooter) applyFooter(d *Document, pdf *gofpdf.Fpdf) error {
	if err := defaults.Set(hf); err != nil {
		return err
	}

	if !hf.UseCustomFunc {
		pdf.SetFooterFunc(func() {
			currentY := pdf.GetY()
			currentX := pdf.GetX()

			pdf.SetTopMargin(HeaderMarginTop)
			pdf.SetY(287 - HeaderMarginTop)
			pdf.SetFont("deja", "", hf.FontSize)
			_, lineHt := pdf.GetFontSize()
			html := pdf.HTMLBasicNew()
			html.Write(lineHt, hf.Text)
			if hf.Pagination {
				pdf.AliasNbPages("")
				pdf.SetY(287 - HeaderMarginTop - 8)
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
