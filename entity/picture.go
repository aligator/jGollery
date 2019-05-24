package entity

import (
	"code.gitea.io/gitea/modules/log"
	"github.com/pkg/errors"
	"jGollery/data"
)

// interface for something which can provide names of children
type ChildNameProvider interface {
	// get the parent's name
	Name() string
	// should return the names of all children
	GetList() ([]string, error)
	// should return the full path for a child or "" if the child doesn't exist
	Get(name string) (string, error)
}

// represents a folder of picture-files
type PictureFiles struct {
	Path string
}

func (p *PictureFiles) Name() string {
	return p.Path
}

func (p *PictureFiles) GetList() ([]string, error) {
	if f, err := data.Open(p.Path); err == nil {
		defer f.Close()
		return f.Pictures()
	} else {
		log.Info("File could not be loaded.", p.Path, err)
		return []string{}, err
	}
}

func (p *PictureFiles) Get(name string) (string, error) {
	fullPath := p.Path + "/" + name
	if f, err := data.Open(fullPath); err == nil {
		defer f.Close()
		if f.IsPicture() {
			return fullPath, nil
		}
		return "", errors.New("file is not a picture " + p.Path)
	} else {
		log.Info("File could not be loaded.", p.Path, err)
		return "", errors.WithMessage(err, "file could not be loaded "+p.Path)
	}
}
