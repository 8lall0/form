package form

import (
	"bytes"
	"io"
	"text/template"

	"github.com/8lall0/form/field"
)

type Form struct {
	Action       string
	Method       string
	Id           string
	IsMultiPart  bool
	Fields       []field.Field
	TemplatePath string
}

func (f *Form) Render(w io.Writer) error {
	var fieldTpl bytes.Buffer
	for _, field := range f.Fields {
		err := field.Render(f.TemplatePath, &fieldTpl)
		if err != nil {
			return err
		}
	}

	tpl, err := template.ParseFiles(f.TemplatePath + "form.html")
	if err != nil {
		return err
	}

	err = tpl.Execute(w, struct {
		Form *Form
		Html string
	}{Form: f, Html: fieldTpl.String()})
	if err != nil {
		return err
	}

	return nil
}
