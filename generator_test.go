package invoice_generator

import (
	"testing"
)

func TestNew(t *testing.T) {
	doc, _ := New(Invoice, &Options{
		TextTypeInvoice: "使用料支払報告書",
		AutoPrint:       true,
	})

	doc.SetFooter(&HeaderFooter{
		Text:       "<center>ご不明な点がございましたら、 DKT株式会社 (contact@postprime.com) までお問い合わせください。</center>",
		Pagination: true,
	})

	doc.SetVersion("IV202107001")

	doc.SetDescription("対象期間：2021年5月1日～2021年5月31日")

	doc.SetDate("2021/07/10")

	doc.SetCompany(&Contact{
		Name:    "DKT株式会社",
		Address: &Address{Country: "日本", City: "東京都", Address: "港区西新橋1-22-5", Address2: "新橋TSビル6階", PostalCode: "〒105-0003"},
		Email:   "contact@postprime.com",
	})

	doc.SetCustomer(&Contact{Name: "安部　慎之介 様"})

	doc.AppendItem(&Item{Name: "対象期間のDKTの売上", Total: 900000, Tax: &Tax{Amount: 90000}})

	doc.SetAfterCommission(&AfterCommission{Amount: 630000})
	doc.SetConsumptionTax(&ConsumptionTax{Amount: 63000})
	doc.SetWithholdingTax(&WithholdingTax{Amount: -64323})
	doc.SetPaymentFree(&PaymentFree{Amount: -35244})
	doc.SetPaidAmount(&PaidAmount{Amount: 593433})

	pdf, err := doc.Build()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = pdf.OutputFileAndClose("payment_details_example.pdf")

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestNew1(t *testing.T) {
	value := formatAmount(100000)
	println(value)
}
