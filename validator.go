/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import "gopkg.in/go-playground/validator.v9"

func (d *Document) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
