/*
Package invoice_generator

Copyright 2021 @chunvv. All Rights Reserved.
Created by Trung Vu on AD 2021/07/10.
*/
package invoice_generator

import (
	"github.com/creasty/defaults"
	"github.com/jung-kurt/gofpdf"
)

func New(docType string, options *Options) (*Document, error) {
	if err := defaults.Set(options); err != nil {
		return nil, err
	}

	doc := &Document{
		Options: options, Type: docType,
	}

	doc.pdf = gofpdf.New("P", "mm", "A4", "")

	return doc, nil
}
