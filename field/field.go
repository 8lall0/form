package field

import (
	"html/template"
	"io"
)

type Field interface {
	Render(string, io.Writer) error
}

func render[T Input | DataList | Button | Select | Checkbox | RadioGroup | TextArea](f *T, tmpl string, w io.Writer) error {
	tpl, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	err = tpl.Execute(w, f)
	if err != nil {
		return err
	}

	return nil
}
