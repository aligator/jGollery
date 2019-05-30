package controller

import (
	"errors"
	"jGollery/data"
	"jGollery/entity"
	"log"
	"net/http"
)

type WebController struct {
	staticPath  string
	galleryPath string
}

func NewWebController(staticPath string, galleryPath string) *WebController {
	wc := WebController{
		staticPath,
		galleryPath,
	}
	err := wc.init()
	if err != nil {
		log.Fatal(err)
	}

	return &wc
}

func (wc *WebController) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (wc *WebController) init() error {
	if err := wc.setupStatic(wc.staticPath); err != nil {
		return err
	}

	if err := wc.setupDynamic(); err != nil {
		return err
	}

	return nil
}

func (wc *WebController) setupStatic(path string) error {
	if wc.check(path) && wc.check(wc.staticPath) {
		fs := http.FileServer(http.Dir(path))
		fullPath := "/" + path + "/"

		http.Handle(fullPath, http.StripPrefix(fullPath, fs))
		return nil
	}
	return errors.New(path + " is not allowed")
}

func (wc *WebController) setupDynamic() error {

	// load all galleries and serve them
	path := wc.staticPath + "/" + wc.galleryPath

	file, err := data.Open(path)
	defer file.Close()

	if err != nil {
		return err
	}

	files, err := file.Readdir(0)
	if err != nil {
		return err
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {

			// serve gallery
			if wc.check(wc.galleryPath) {
				pics := entity.NewPicFiles(path, fileInfo.Name())
				gallery, err := entity.NewGallery(pics)

				if err == nil {
					http.Handle("/"+fileInfo.Name()+"/", gallery)
				} else {
					return errors.New("could not set up gallery")
				}
			} else {
				return errors.New(wc.galleryPath + " is not allowed")
			}
		}
	}

	return nil
}

func (wc *WebController) check(path string) bool {
	// TODO
	return true
}
