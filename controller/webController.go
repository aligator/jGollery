package controller

import (
	"errors"
	"jGollery/entity"
	"log"
	"net/http"
)

type WebController struct {
	galleryPath string
	webPath     string
}

const staticPath = "static"

func NewWebController(galleryPath string, webPath string) *WebController {
	wc := WebController{
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (wc *WebController) init() error {
	if err := wc.setupStatic(staticPath); err != nil {
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
		http.Handle("/"+path+"/", http.StripPrefix("/"+path+"/", fs))
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
