package entity

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type Templated interface {
	RenderTemplate(writer http.ResponseWriter)
	WriteTemplate(writer io.Writer) error
}

type TemplateData interface{}

type Template struct {
	templateName string
}

var templates = template.New("none")

func (comp *Template) MustOrInternalServerError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, "internal error", http.StatusInternalServerError)

		log.Fatal(err.Error())
	}
}

func (comp *Template) WriteTemplate(writer io.Writer) error {
	return comp.writeTemplate(writer, nil)
}

func (comp *Template) RenderTemplate(writer http.ResponseWriter) {
	comp.renderTemplate(writer, nil)
}

func (comp *Template) renderTemplate(writer http.ResponseWriter, compData TemplateData) {
	err := comp.writeTemplate(writer, compData)
	comp.MustOrInternalServerError(writer, err)
}

func (comp *Template) writeTemplate(writer io.Writer, compData TemplateData) error {
	err := templates.ExecuteTemplate(writer, comp.templateName, compData)
	return err
}

// Todo: combine the new Methods, as there is duplicated code
func NewTemplate(path string, filename string) *Template {
	fullPath := path + "/" + filename
	if templ := templates.Lookup(filename); templ != nil {
		return &Template{
			templateName: filename,
		}
	} else {
		template.Must(templates.ParseFiles(fullPath))
	}

	return &Template{
		templateName: filename,
	}
}

func NewDirectTemplate(name string, templStr string) *Template {
	if templ := templates.Lookup(name); templ != nil {
		return &Template{
			templateName: name,
		}
	} else {
		template.Must(templates.New(name).Parse(templStr))
	}

	return &Template{
		templateName: name,
	}
}
