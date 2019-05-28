package entity

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
)

type Gallery struct {
	Template
	Data []string
}

func (g *Gallery) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := g.WriteTemplate(writer)
	if err != nil {
		log.Println("could not serve gallery", err)
		http.Error(writer, "could not serve gallery", http.StatusInternalServerError)
	}
}

func (g *Gallery) WriteTemplate(writer io.Writer) error {
	return g.writeTemplate(writer, g.Data)
}

func (g *Gallery) RenderTemplate(writer http.ResponseWriter) {
	g.renderTemplate(writer, g.Data)
}

func NewGallery(pictures PathGroup) (*Gallery, error) {
	pics, err := pictures.GetList()
	if err != nil {
		return nil, errors.Errorf("Could not load gallery %s", pictures.Name())
	}

	picPaths := make([]string, len(pics))

	for _, pic := range pics {
		if picPath, err := pictures.Get(pic); err == nil {
			picPaths = append(picPaths, "/"+picPath)
		} else {
			return nil, errors.Errorf("%s not found\n", pic)
		}
	}

	component := &Gallery{
		Template: *NewTemplate("template", "gallery.html"),
		Data:     picPaths,
	}

	return component, nil
}
