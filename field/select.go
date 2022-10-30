package field

import "io"

type SelectOption struct {
	Value    string
	Label    string
	Selected bool
}

type Select struct {
	Id       string
	Name     string
	Label    string
	Required bool
	Size     int
	Multiple bool
	Options  []SelectOption
}

func (s *Select) Render(tplPath string, w io.Writer) error {
	if err := render(s, tplPath+tplSelect, w); err != nil {
		return err
	}
	return nil
}
