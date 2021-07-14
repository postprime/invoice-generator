/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"math/rand"
	"time"
)

var Letters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func encodeString(str string) string {
	return str
}

func (d *Document) typeAsString() string {
	if d.Type == Invoice {
		return d.Options.TextTypeInvoice
	}
	if d.Type == Receipt {
		return d.Options.TextTypeInvoice
	}
	return d.Options.TextTypeInvoice
}

func formatAmount(v int) string {
	cur := currency.MustParseISO("JPY")
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(v, number.Scale(scale))
	p := message.NewPrinter(language.Japanese)
	return p.Sprintf("%v%v", currency.Symbol(cur), dec)
}

func GenerateInvoiceNumber() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 8)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(b) + "-001"
}
