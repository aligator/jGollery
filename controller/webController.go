package controller

import (
	"errors"
	"jGollery/entity"
	"log"
	"net/http"
)

type WebController struct {
	staticPath  string
	galleryPath string
	webPath     string
}

func NewWebController(staticPath string, galleryPath string, webPath string) *WebController {
	wc := WebController{
		staticPath,
		galleryPath,
		webPath,
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
	if wc.check(path) {
		fs := http.FileServer(http.Dir(path))
		fullPath := "/" + wc.staticPath + "/" + path + "/"

		http.Handle(fullPath, http.StripPrefix(fullPath, fs))
		return nil
	}
	return errors.New(path + " is not allowed")
}

func (wc *WebController) setupDynamic() error {
	if wc.check(wc.galleryPath) {
		pics := entity.PictureFiles{Path: wc.galleryPath}
		g, err := entity.NewGallery(&pics)

		if err == nil {
			http.Handle("/"+wc.webPath+"/", g)
		} else {
			return errors.New("could not set up gallery")
		}
	} else {
		return errors.New(wc.galleryPath + " is not allowed")
	}

	return nil
}

func (wc *WebController) check(path string) bool {
	// TODO
	return true
}
