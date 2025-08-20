package xmltmpl

import (
	"discaunter/domain"
)

type templateString struct {
	app domain.Apper
}

func NewTemplate(app domain.Apper) *templateString {
	return &templateString{
		app: app,
	}
}
