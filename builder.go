/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

func (d *Document) SetType(docType string) *Document {
	d.Type = docType
	return d
}

func (d *Document) SetHeader(header *HeaderFooter) *Document {
	d.Header = header
	return d
}

func (d *Document) SetFooter(footer *HeaderFooter) *Document {
	d.Footer = footer
	return d
}

func (d *Document) SetVersion(version string) *Document {
	d.Number = version
	return d
}

func (d *Document) SetDescription(desc string) *Document {
	d.Description = desc
	return d
}

func (d *Document) SetNotes(notes string) *Document {
	d.Notes = notes
	return d
}

func (d *Document) SetCompany(company *Contact) *Document {
	d.Company = company
	return d
}

func (d *Document) SetCustomer(customer *Contact) *Document {
	d.Customer = customer
	return d
}

func (d *Document) AppendItem(item *Item) *Document {
	d.Items = append(d.Items, item)
	return d
}

func (d *Document) SetDate(date string) *Document {
	d.Date = date
	return d
}

func (d *Document) SetPaymentTerm(term string) *Document {
	d.PaymentTerm = term
	return d
}

func (d *Document) SetAfterCommission(afterCommission *AfterCommission) *Document {
	d.AfterCommission = afterCommission
	return d
}

func (d *Document) SetConsumptionTax(consumptionTax *ConsumptionTax) *Document {
	d.ConsumptionTax = consumptionTax
	return d
}

func (d *Document) SetWithholdingTax(withholdingTax *WithholdingTax) *Document {
	d.WithholdingTax = withholdingTax
	return d
}

func (d *Document) SetPaymentFree(paymentFee *PaymentFree) *Document {
	d.PaymentFree = paymentFee
	return d
}

func (d *Document) SetPaidAmount(paidAmount *PaidAmount) *Document {
	d.PaidAmount = paidAmount
	return d
}

func (d *Document) SetPaidAmount1(paidAmount *PaidAmount) *Document {
	d.PaidAmount1 = paidAmount
	return d
}
func (d *Document) SetPaidAmount2(paidAmount *PaidAmount) *Document {
	d.PaidAmount2 = paidAmount
	return d
}
