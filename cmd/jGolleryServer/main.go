package main

import (
	"flag"
	. "jGollery/controller"
)

func main() {
	addr := flag.String("addr", ":8080", `The address the application servers listens to`)
	flag.Parse()

	webController := NewWebController("static", "gallery")
	webController.Run(*addr)
}
