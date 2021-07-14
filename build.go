/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

func (d *Document) Build() (*gofpdf.Fpdf, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}

	d.pdf.SetMargins(BaseMargin, BaseMarginTop, BaseMargin)
	d.pdf.SetXY(10, 10)
	d.pdf.SetTextColor(BaseTextColor[0], BaseTextColor[1], BaseTextColor[2])

	if d.Header != nil {
		err = d.Header.applyHeader(d, d.pdf)

		if err != nil {
			return nil, err
		}
	}

	if d.Footer != nil {
		err = d.Footer.applyFooter(d, d.pdf)

		if err != nil {
			return nil, err
		}
	}

	d.pdf.AddPage()
	d.pdf.AddUTF8Font("deja", "", "./fonts/japanese-font.ttf")
	d.pdf.SetFont("deja", "", 12)

	d.appendTitle(d.pdf)
	d.appendMetas(d.pdf)

	companyBottom := d.Company.appendCompanyContactToDoc(d.pdf)
	customerBottom := d.Customer.appendCustomerContactToDoc(d.pdf)
	if customerBottom > companyBottom {
		d.pdf.SetXY(10, customerBottom)
	} else {
		d.pdf.SetXY(10, companyBottom)
	}
	d.pdf.SetY(d.pdf.GetY() + 10)

	d.appendDescription(d.pdf)

	//d.appendItems(d.pdf)
	d.appendNotes(d.pdf)
	d.appendTotal(d.pdf)
	if d.Options.AutoPrint {
		d.pdf.SetJavascript("print(true);")
	}

	return d.pdf, nil
}

func (d *Document) appendTitle(pdf *gofpdf.Fpdf) {
	title := d.typeAsString()
	pdf.SetXY(120, BaseMarginTop)
	pdf.SetFillColor(DarkBgColor[0], DarkBgColor[1], DarkBgColor[2])
	pdf.Rect(120, BaseMarginTop, 80, 10, "F")
	d.pdf.SetFont("deja", "", 20)
	pdf.CellFormat(80, 10, title, "0", 0, "C", false, 0, "")
}

func (d *Document) appendMetas(pdf *gofpdf.Fpdf) {
	if len(d.Number) > 0 {
		versionString := fmt.Sprintf("%s: %s", d.Options.TextNumberTitle, d.Number)
		pdf.SetXY(120, BaseMarginTop+15)
		pdf.SetFont("deja", "", 8)
		pdf.CellFormat(80, 4, versionString, "0", 0, "R", false, 0, "")
	}

	date := time.Now().Format("02/01/2006")
	if len(d.Date) > 0 {
		date = d.Date
	}
	dateString := fmt.Sprintf("%s: %s", d.Options.TextDateTitle, date)
	pdf.SetXY(120, BaseMarginTop+19)
	pdf.SetFont("deja", "", 8)
	pdf.CellFormat(80, 4, encodeString(dateString), "0", 0, "R", false, 0, "")
}

func (d *Document) appendDescription(pdf *gofpdf.Fpdf) {
	if len(d.Description) > 0 {
		pdf.SetY(pdf.GetY() + 10)
		pdf.SetFont("deja", "", 10)
		pdf.MultiCell(190, 5, encodeString(d.Description), "B", "L", false)
	}
}

func (d *Document) drawsTableTitles(pdf *gofpdf.Fpdf) {
	pdf.SetX(10)
	pdf.SetY(pdf.GetY() + 5)
	pdf.SetFont("deja", "", 8)
	pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
	pdf.Rect(10, pdf.GetY(), 190, 6, "F")

	pdf.SetX(ItemColNameOffset)
	pdf.CellFormat(ItemColUnitPriceOffset-ItemColNameOffset, 6, encodeString(d.Options.TextItemsNameTitle), "0", 0, "", false, 0, "")

	pdf.SetX(ItemColTotalHTOffset)
	pdf.CellFormat(ItemColTaxOffset-ItemColTotalHTOffset, 6, encodeString(d.Options.TextItemsTotalTTCTitle), "0", 0, "", false, 0, "")

	pdf.SetX(ItemColTotalTTCOffset)
	pdf.CellFormat(190-ItemColTotalTTCOffset, 6, encodeString(d.Options.TextItemsTotalConsumptionTaxTitle), "0", 0, "", false, 0, "")
}

func (d *Document) appendItems(pdf *gofpdf.Fpdf) {
	d.drawsTableTitles(pdf)

	pdf.SetX(10)
	pdf.SetY(pdf.GetY() + 8)
	pdf.SetFont("deja", "", 8)

	for i := 0; i < len(d.Items); i++ {
		item := d.Items[i]

		item.appendColTo(d.Options, pdf)

		if pdf.GetY() > MaxPageHeight {
			pdf.AddPage()
			d.drawsTableTitles(pdf)
			pdf.SetFont("deja", "", 8)
		}

		pdf.SetX(10)
		pdf.SetY(pdf.GetY() + 6)
	}
}

func (d *Document) appendNotes(pdf *gofpdf.Fpdf) {
	if len(d.Notes) == 0 {
		return
	}

	currentY := pdf.GetY()

	pdf.SetFont("deja", "", 9)
	pdf.SetX(BaseMargin)
	pdf.SetRightMargin(100)
	pdf.SetY(currentY + 10)

	_, lineHt := pdf.GetFontSize()
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, encodeString(d.Notes))

	pdf.SetRightMargin(BaseMargin)
	pdf.SetY(currentY)
}

func (d *Document) appendTotal(pdf *gofpdf.Fpdf) {
	total, _ := decimal.NewFromString("0")

	for _, item := range d.Items {
		temp, _ := decimal.NewFromString(strconv.Itoa(item.Total))
		total = total.Add(temp)
	}

	//// Tax
	//totalTax := decimal.NewFromFloat(0)
	//if d.Discount == nil {
	//	for _, item := range d.Items {
	//		totalTax = totalTax.Add(item.taxWithDiscount())
	//	}
	//} else {
	//	discountType, discountAmount := d.Discount.getDiscount()
	//	discountPercent := discountAmount
	//	if discountType == "amount" {
	//		// Get percent from total discounted
	//		discountPercent = discountAmount.Mul(decimal.NewFromFloat(100)).Div(totalWithDiscount)
	//	}
	//
	//	for _, item := range d.Items {
	//		if item.Tax != nil {
	//			taxType, taxAmount := item.Tax.getTax()
	//			if taxType == "amount" {
	//				// If tax type is amount, juste add amount to tax
	//				totalTax = totalTax.Add(taxAmount)
	//			} else {
	//				// Else, remove doc discount % from item total without tax and item discount
	//				itemTotal := item.totalWithoutTaxAndWithDiscount()
	//				toSub := discountPercent.Mul(itemTotal).Div(decimal.NewFromFloat(100))
	//				itemTotalDiscounted := itemTotal.Sub(toSub)
	//
	//				// Then recompute tax on itemTotalDiscounted
	//				itemTaxDiscounted := taxAmount.Mul(itemTotalDiscounted).Div(decimal.NewFromFloat(100))
	//
	//				totalTax = totalTax.Add(itemTaxDiscounted)
	//			}
	//		}
	//	}
	//}
	//
	//// finalTotal
	//totalWithTax := total.Add(totalTax)
	//if d.Discount != nil {
	//	totalWithTax = totalWithDiscount.Add(totalTax)
	//}

	pdf.SetY(pdf.GetY() + 5)
	// Draw after commission
	pdf.SetX(162)
	pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
	pdf.Rect(160, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(40, 10, formatAmount(d.AfterCommission.Amount), "0", 0, "L", false, 0, "")

	pdf.SetX(120)
	pdf.Rect(120, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(38, 10, encodeString(d.Options.TextAfterCommissionTotal), "0", 0, "R", false, 0, "")
	pdf.SetY(pdf.GetY() + 11)

	if d.ConsumptionTax != nil {
		// Draw consumption tax
		pdf.SetX(162)
		pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
		pdf.Rect(160, pdf.GetY(), 40, 10, "F")
		pdf.CellFormat(40, 10, formatAmount(d.ConsumptionTax.Amount), "0", 0, "L", false, 0, "")

		pdf.SetX(120)
		pdf.Rect(120, pdf.GetY(), 40, 10, "F")
		pdf.CellFormat(38, 10, encodeString(d.Options.TextConsumptionTaxTotal), "0", 0, "R", false, 0, "")
		pdf.SetY(pdf.GetY() + 15)
	}

	// Draw withholding tax
	pdf.SetX(162)
	pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
	pdf.Rect(160, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(40, 10, "("+formatAmount(d.WithholdingTax.Amount)+")", "0", 0, "L", false, 0, "")

	pdf.SetX(120)
	pdf.Rect(120, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(38, 10, encodeString(d.Options.TextWithholdingTaxTotal), "0", 0, "R", false, 0, "")
	pdf.SetY(pdf.GetY() + 11)

	// Draw payment free
	pdf.SetX(162)
	pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
	pdf.Rect(160, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(40, 10, "("+formatAmount(d.PaymentFree.Amount)+")", "0", 0, "L", false, 0, "")

	pdf.SetX(120)
	pdf.Rect(120, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(38, 10, encodeString(d.Options.TextPaymentTotal), "0", 0, "R", false, 0, "")
	pdf.SetY(pdf.GetY() + 11)

	// Draw paid amount
	pdf.SetX(162)
	pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
	pdf.Rect(160, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(40, 10, formatAmount(d.PaidAmount.Amount), "0", 0, "L", false, 0, "")

	pdf.SetX(120)
	pdf.SetFillColor(GreyBgColor[0], GreyBgColor[1], GreyBgColor[2])
	pdf.Rect(120, pdf.GetY(), 40, 10, "F")
	pdf.CellFormat(38, 10, encodeString(d.Options.TextTotalTotal), "0", 0, "R", false, 0, "")
}
