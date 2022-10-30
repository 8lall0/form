package form

import (
	"bytes"
	"io"
	"text/template"

	"github.com/8lall0/form/field"
)

const (
	tplForm = "form.html"
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
	for _, el := range f.Fields {
		err := el.Render(f.TemplatePath, &fieldTpl)
		if err != nil {
			return err
		}
	}

	tpl, err := template.ParseFiles(f.TemplatePath + tplForm)
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
