package main

import (
	"fmt"
	"jGollery"
)

func main() {
	f, err := jGollery.Open("gallery/demo")
	must(err)

	pics, err := f.Pictures()
	must(err)

	for _, pic := range *pics {
		fmt.Println(pic.Name())
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
