package xmltmpl

import (
	"bytes"
	"encoding/xml"
	"text/template"
)

var funcMapText = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
	"escape": func(s string) string {
		var sh bytes.Buffer
		xml.Escape(&sh, []byte(s))
		return sh.String()
	},
}
