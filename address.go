/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

type Address struct {
	Address    string `json:"address,omitempty" validate:"required"`
	Address2   string `json:"address_2,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
}

func (a *Address) ToString() string {
	var addrString = a.PostalCode
	if len(a.City) > 0 {
		addrString += "\n"
		addrString += a.City
	}

	if len(a.Address) > 0 {
		addrString += a.Address
	} else {
		addrString += "\n"
	}

	if len(a.Address2) > 0 {
		addrString += "\n"
		addrString += a.Address2
	}

	return addrString
}
