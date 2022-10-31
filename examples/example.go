package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/8lall0/form"
	"github.com/8lall0/form/field"
)

func main() {

	input := &field.Input{
		Type:        "text",
		Id:          "text",
		Name:        "Text",
		Placeholder: "Text",
		Label:       "Text Field",
	}

	datalist := &field.DataList{
		Id:      "datalist",
		Name:    "Datalist",
		Label:   "Datalist field",
		Options: []string{"data one", "data two"},
	}
	button := &field.Button{
		Label: "button",
	}

	options := make([]field.SelectOption, 0)
	options = append(options, field.SelectOption{Value: "0", Label: "value zero"})
	options = append(options, field.SelectOption{Value: "1", Label: "value one"})
	sel := &field.Select{
		Id:      "select",
		Name:    "select",
		Label:   "Select Field",
		Options: options,
	}

	val1 := field.RadioValue{
		Id:    "value one",
		Value: "value one",
		Label: "value one",
	}
	val2 := field.RadioValue{
		Id:    "value two",
		Value: "value two",
		Label: "value two",
	}
	rGroup := &field.RadioGroup{
		Name:     "radiogroup",
		Label:    "Radiogroup Field",
		Required: true,
		Values:   []field.RadioValue{val1, val2},
	}
	check := &field.Checkbox{
		Name:     "checkbox",
		Id:       "checkbox",
		Label:    "Checkbox",
		Value:    "checkbox",
		Required: true,
	}
	txtArea := &field.TextArea{
		Id:    "textarea",
		Name:  "textarea",
		Label: "Textarea Field",
	}

	exampleForm := form.Form{
		Method:       "GET",
		Action:       "#",
		TemplatePath: "template/default/",
	}

	exampleForm.Fields = []field.Field{input, datalist, sel, check, rGroup, txtArea, button}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := fmt.Fprintln(w, "<html><body><head><style>body{font-family: sans-serif;}div + div {margin-top: 20px;}</style></head>")
		if err != nil {
			log.Println(err)
		}
		err = exampleForm.Render(w)
		if err != nil {
			log.Println(err)
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
