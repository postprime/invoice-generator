/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import (
	"bytes"
	b64 "encoding/base64"
	"github.com/jung-kurt/gofpdf"
	"image"
)

type Contact struct {
	Name    string   `json:"name,omitempty" validate:"required,min=1,max=256"`
	Logo    *[]byte  `json:"logo,omitempty"`
	Address *Address `json:"address,omitempty"`
	Email   string   `json:"email,omitempty"`
}

func (c *Contact) appendContactTODoc(x float64, y float64, fill bool, logoAlign string, pdf *gofpdf.Fpdf) float64 {
	pdf.SetXY(x, y)

	if c.Logo != nil {
		fileName := b64.StdEncoding.EncodeToString([]byte(c.Name))
		ioReader := bytes.NewReader(*c.Logo)
		_, format, _ := image.DecodeConfig(bytes.NewReader(*c.Logo))
		imageInfo := pdf.RegisterImageOptionsReader(fileName, gofpdf.ImageOptions{
			ImageType: format,
		}, ioReader)

		if imageInfo != nil {
			var imageOpt gofpdf.ImageOptions
			imageOpt.ImageType = format

			pdf.ImageOptions(fileName, pdf.GetX(), y, 0, 30, false, imageOpt, 0, "")

			pdf.SetY(y + 30)
		}
	}

	if fill {
		pdf.SetFillColor(255, 255, 255)
	} else {
		pdf.SetFillColor(255, 255, 255)
	}

	pdf.SetX(x)

	pdf.Rect(x, pdf.GetY(), 70, 8, "F")
	pdf.SetFont("deja", "", 14)
	pdf.Cell(40, 8, c.Name)
	pdf.SetFont("deja", "", 10)

	if c.Address != nil {
		var addrRectHeight float64 = 17

		if len(c.Address.Address2) > 0 {
			addrRectHeight = addrRectHeight + 5
		}

		if len(c.Address.Country) == 0 {
			addrRectHeight = addrRectHeight - 5
		}

		pdf.Rect(x, pdf.GetY()+9, 70, addrRectHeight, "F")

		pdf.SetFont("deja", "", 10)
		pdf.SetXY(x, pdf.GetY()+10)
		pdf.MultiCell(70, 5, c.Email, "0", "L", false)
		pdf.MultiCell(70, 5, c.Address.ToString(), "0", "L", false)
		pdf.SetXY(x, pdf.GetY()+10)
	}

	return pdf.GetY()
}

func (c *Contact) appendCompanyContactToDoc(pdf *gofpdf.Fpdf) float64 {
	x, y, _, _ := pdf.GetMargins()
	return c.appendContactTODoc(x, y, true, "L", pdf)
}

func (c *Contact) appendCustomerContactToDoc(pdf *gofpdf.Fpdf) float64 {
	x, _, _, _ := pdf.GetMargins()
	return c.appendContactTODoc(x, pdf.GetY(), true, "L", pdf)
}
