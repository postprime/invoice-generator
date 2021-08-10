/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/11.
*/
package invoice_generator

type AfterCommission struct {
	Amount            int  `json:"amount,omitempty"`
	ConsumptionTax    int  `json:"consumption_tax,omitempty"`
	IsDomesticCreator bool `json:"is_domestic_creator,omitempty"`
}
