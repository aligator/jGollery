package main

import (
	"jGollery/entity"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	pics := entity.PictureFiles{Path: "gallery/demo"}
	g, err := entity.NewGallery(&pics)
	Must(err)
	g.RenderTemplate(w)
}
