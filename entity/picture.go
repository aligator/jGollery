package entity

import (
	"errors"
	"jGollery/data"
	"log"
	"regexp"
)

var galleryRegexp = regexp.MustCompile("^[a-zA-Z0-9]+$")

// interface for something which groups children beneath a path
type PathGroup interface {
	// get the parent's name
	Name() string
	// returns the names of all children
	GetList() ([]string, error)
	// returns the full path for a child (parent name + child name) or error if the child doesn't exist
	Get(name string) (string, error)
}

// represents a folder of picture-files
type PictureFiles struct {
	name string
	path string
}

func NewPicFiles(folder string, name string) *PictureFiles {
	pf := PictureFiles{
		name: name,
		path: folder + "/" + name,
	}

	return &pf
}

func (p *PictureFiles) Name() string {
	return p.name
}

func (p *PictureFiles) GetList() ([]string, error) {
	if f, err := data.Open(p.path); err == nil {
		defer f.Close()
		return f.Pictures()
	} else {
		log.Println("file could not be loaded.", p.path, err)
		return []string{}, err
	}
}

func (p *PictureFiles) Get(name string) (string, error) {
	fullPath := p.path + "/" + name
	if f, err := data.Open(fullPath); err == nil {
		defer f.Close()
		if f.IsPicture() {
			return fullPath, nil
		}
		return "", errors.New("file is not a picture " + p.path)
	} else {
		log.Println("file could not be loaded", p.path, err)
		return "", err
	}
}
