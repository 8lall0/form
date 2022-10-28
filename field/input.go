package field

import (
	"io"
)

type Input struct {
	Type         string
	Id           string
	Name         string
	Placeholder  string
	Value        string
	Label        string
	Required     bool
	ReadOnly     bool
	Disabled     bool
	Size         int
	MaxLength    int
	Min          string
	Max          string
	Pattern      string
	AutoFocus    bool
	AutoComplete bool
}

func (i *Input) Render(tplPath string, w io.Writer) error {
	// Qui puoi controllare i valori
	if err := render(i, tplPath+"input.html", w); err != nil {
		return err
	}

	return nil
}
