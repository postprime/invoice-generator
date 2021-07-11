/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/

package invoice_generator

const (
	Invoice         string  = "INVOICE"
	Receipt         string  = "RECEIPT"
	BaseMargin      float64 = 10
	BaseMarginTop   float64 = 20
	HeaderMarginTop float64 = 5
	MaxPageHeight   float64 = 260
)

const (
	ItemColNameOffset      float64 = 10
	ItemColUnitPriceOffset float64 = 80
	ItemColQuantityOffset  float64 = 103
	ItemColTotalHTOffset   float64 = 113
	ItemColDiscountOffset  float64 = 140
	ItemColTaxOffset       float64 = 157
	ItemColTotalTTCOffset  float64 = 175
)

var (
	BaseTextFontSize  float64 = 8
	SmallTextFontSize float64 = 7
	LargeTextFontSize float64 = 10
	BaseTextColor             = []int{35, 35, 35}
	GreyTextColor             = []int{82, 82, 82}
	GreyBgColor               = []int{232, 232, 232}
	DarkBgColor               = []int{212, 212, 212}
)
