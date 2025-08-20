package xmltmpl

import (
	_ "embed"
	"fmt"
)

//go:embed tmplXml.xml
var tmplXml string

func (tt *templateString) StringXML(model interface{}) (bts []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic templateString %v", r)
		}
	}()

	tmplName := "xml"
	// вызов шаблона в него передаем имя шаблона как имя файла шаблона
	if result, err := tt.tmplMustText(tmplXml, tmplName, model, nil); err != nil {
		return bts, err
	} else {
		return result, err
	}
}
