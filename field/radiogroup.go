package field

import "io"

type RadioValue struct {
	Id      string
	Value   string
	Label   string
	Checked bool
}

type RadioGroup struct {
	Name     string
	Label    string
	Required bool
	Values   []RadioValue
}

func (r *RadioGroup) Render(tplPath string, w io.Writer) error {
	if err := render(r, tplPath+tplRadioGroup, w); err != nil {
		return err
	}
	return nil
}
