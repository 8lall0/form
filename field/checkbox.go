package field

import "io"

type Checkbox struct {
	Name     string
	Id       string
	Value    string
	Label    string
	Required bool
}

func (c *Checkbox) Render(tplPath string, w io.Writer) error {
	if err := render(c, tplPath+"checkbox.html", w); err != nil {
		return err
	}
	return nil
}
