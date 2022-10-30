package field

import "io"

type TextArea struct {
	Id          string
	Name        string
	Placeholder string
	Value       string
	Label       string
	Rows        int
	Cols        int
	Required    bool
}

func (t *TextArea) Render(tplPath string, w io.Writer) error {
	if err := render(t, tplPath+tplTextarea, w); err != nil {
		return err
	}
	return nil
}
