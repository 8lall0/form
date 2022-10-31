package field

import (
	"io"
)

type DataList struct {
	Id       string
	Name     string
	Label    string
	Type     string
	Required bool
	Size     int
	Multiple bool
	Options  []string
}

func (d *DataList) Render(tplPath string, w io.Writer) error {
	if err := render(d, tplPath+tplDatalist, w); err != nil {
		return err
	}
	return nil
}
