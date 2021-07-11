/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

type Tax struct {
	Percent string `json:"percent,omitempty"`
	Amount  int    `json:"amount,omitempty"`
}
