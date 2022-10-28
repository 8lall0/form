package field

import (
	"io"
)

type DataList struct {
	Id       string
	Name     string
	Label    string
	Required bool
	Size     int
	Multiple bool
	Options  []string
	Validate func() bool
}

func (d *DataList) Render(tplPath string, w io.Writer) error {
	if err := render(d, tplPath+"datalist.html", w); err != nil {
		return err
	}
	return nil
}
