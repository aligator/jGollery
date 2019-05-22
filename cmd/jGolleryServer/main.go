package main

import (
	"fmt"
	"jGollery/entity"
)

func main() {

	p := entity.PictureFiles{Path: "gallery/demo"}

	pics, err := p.GetList()
	Must(err)

	for _, pic := range pics {
		fmt.Println(p.Get(pic))
	}

}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
