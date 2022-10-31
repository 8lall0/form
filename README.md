# Form
A simple library to generate HTML forms.

## System Requirements
You need a Go version > 1.19, due to usage of generics.

## Install
`go get github.com/8lall0/form`

## Usage
Create a new `form`, create some `form/field` fields and append them to the `Fields` attribute of `form` in order of appearance.

When in doubt, just see the examples in the `examples/` directory.

## Template
The examples come with some ready to use templates inside `examples/template/default`.

If you want to use a custom template, you can modify the default one or create one from scratch, following the same filenames. You can pass the new template path inside the `TemplatePath` attribute of a `form`.

## Custom fields
You can define custom fields as you want, they only have to satisfy the `form/field.Field` interface, which requires this signature:

`Render(string, io.Writer) error`

## Contributions
Feel free to fork and send pull requests :)

## License
This project is licensed under the MIT License - see the LICENSE.md file for details