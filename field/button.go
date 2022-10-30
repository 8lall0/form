package field

import "io"

type Button struct {
	Type  string
	Id    string
	Value string
	Label string
	Name  string
}

func (b *Button) Render(tplPath string, w io.Writer) error {
	if err := render(b, tplPath+tplButton, w); err != nil {
		return err
	}
	return nil
}
