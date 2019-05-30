package controller

import (
	"errors"
	"jGollery/data"
	"jGollery/entity"
	"log"
	"net/http"
	"regexp"
)

type WebController struct {
	staticPath  string
	galleryPath string
}

// Controller for all webservices of jGollery
// staticPath is the folder where all static files are located. Such as the gallery-images or stylesheets.
// galleryPath is the folder inside of staticPath, where the folders for the galleries are located.
func NewWebController(staticPath string, galleryPath string) *WebController {
	wc := WebController{
		staticPath,
		galleryPath,
	}
	if wc.check(staticPath) && wc.check(galleryPath) {
		err := wc.init()
		if err != nil {
			log.Fatal(err)
		}

		return &wc
	} else {
		log.Fatal("given paths are not valid", staticPath, galleryPath)
		return nil
	}
}

// run the webserver
func (wc *WebController) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}

// setup static fileserver and dynamic webserver for jGollery
func (wc *WebController) init() error {
	if err := wc.setupStatic(wc.staticPath); err != nil {
		return err
	}

	if err := wc.setupDynamic(); err != nil {
		return err
	}

	return nil
}

// setups a fileserver which serves all files inside the given path.
func (wc *WebController) setupStatic(path string) error {
	fs := http.FileServer(http.Dir(path))
	fullPath := "/" + path + "/"

	http.Handle(fullPath, http.StripPrefix(fullPath, fs))
	return nil
}

// setups the webserver which handles gallery requests. Uses the paths provided by the webController.
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

	return nil
}

const pathRegexp = "^[\\/a-zA-Z0-9\\-_ äöü]+$"

var compPathRegexp = regexp.MustCompile(pathRegexp)

// checks if path is valid
func (wc *WebController) check(path string) bool {
	return compPathRegexp.MatchString(path)
}
