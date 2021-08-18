/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

type Options struct {
	AutoPrint bool `json:"auto_print,omitempty"`

	CurrencySymbol    string `default:"円" json:"currency_symbol,omitempty"`
	CurrencyPrecision int    `default:"2" json:"currency_precision,omitempty"`
	CurrencyDecimal   string `default:"." json:"currency_decimal,omitempty"`
	CurrencyThousand  string `default:" " json:"currency_thousand,omitempty"`

	TextTypeInvoice string `default:"請求書" json:"text_type_invoice,omitempty"`
	TextTypeReceipt string `default:"領収書" json:"text_type_delivery_note,omitempty"`

	TextNumberTitle string `default:"ロイヤリティ支払報告書No" json:"text_version_title,omitempty"`
	TextDateTitle   string `default:"発行日" json:"text_date_title,omitempty"`

	TextItemsNameTitle                string `default:"摘要" json:"text_items_name_title,omitempty"`
	TextItemsTotalTTCTitle            string `default:"売上金額" json:"text_items_total_ttc_title,omitempty"`
	TextItemsTotalConsumptionTaxTitle string `default:"消費税額" json:"text_items_total_consumption_tax_title,omitempty"`

	TextAfterCommissionTotal        string `default:"ロイヤリティ(消費税込み)" json:"text_commission_total,omitempty"`
	TextAfterCommissionTotalOversea string `default:"ロイヤリティ" json:"text_commission_total_oversea,omitempty"`
	TextConsumptionTaxTotal         string `default:"消費税額等" json:"text_consumption_tax_total,omitempty"`
	TextPaymentTotal                string `default:"決済手数料" json:"text_payment_total,omitempty"`
	TextWithholdingTaxTotal         string `default:"源泉所得税" json:"text_withholding_tax_total,omitempty"`
	TextTotalTotal                  string `default:"お支払い金額" json:"text_total_total,omitempty"`
	Payout1                         string `default:"第1回支払日" json:"payout_1,omitempty"`
	Payout2                         string `default:"第2回支払日" json:"payout_2,omitempty"`
}
