package invoice_generator

import (
	"testing"
)

func TestNew(t *testing.T) {
	doc, _ := New(Invoice, &Options{
		TextTypeInvoice: "ロイヤリティ支払報告書",
		AutoPrint:       true,
	})

	doc.SetFooter(&HeaderFooter{
		Text: "<center>*ロイヤリティは所得税法上の源泉徴収の対象となる使用料の支払に該当し、所定の源泉徴収税率を乗じて算出しています。</center>\n" +
			"<center>*上記を確認の上、内容に異議がある場合には2週間以内に連絡下さいますようお願い致します。</center>",
		Pagination: true,
	})

	doc.SetVersion(GenerateInvoiceNumber())

	doc.SetDescription("対象期間：2021年5月1日～2021年5月31日")

	doc.SetDate("2021/07/10")

	doc.SetCompany(&Contact{
		Name:    "PostPrime(DKT株式会社)",
		Address: &Address{Country: " ", City: " ", Address: " ", Address2: " ", PostalCode: " "},
	})

	doc.SetCustomer(&Contact{Name: "安部　慎之介 様"})
	doc.SetInvoiceRegistrationNumber("T1234567891234")

	//doc.AppendItem(&Item{Name: "対象期間のDKTの売上", Total: 900000, Tax: &Tax{Amount: 90000}})

	doc.SetAfterCommission(&AfterCommission{Amount: 630000, ConsumptionTax: 64545, IsDomesticCreator: true})
	doc.SetWithholdingTax(&WithholdingTax{Amount: 64323})
	doc.SetPaymentFree(&PaymentFree{Amount: 35244})
	doc.SetPaidAmount(&PaidAmount{Amount: 593433, PayoutDate: "2021年11月30日"})

	pdf, err := doc.Build()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = pdf.OutputFileAndClose("payment_details_example.pdf")

	if err != nil {
		t.Errorf(err.Error())
	}
}
