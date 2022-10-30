package field

import (
	"html/template"
	"io"
)

type Field interface {
	Render(string, io.Writer) error
}

const (
	tplButton     = "button.html"
	tplCheckbox   = "checkbox.html"
	tplDatalist   = "datalist.html"
	tplInput      = "input.html"
	tplRadioGroup = "radiogroup.html"
	tplTextarea   = "textarea.html"
	tplSelect     = "select.html"
)

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
