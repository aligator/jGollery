package main

import . "jGollery/controller"

func main() {
	webController := NewWebController("static/gallery/demo", "demo")
	webController.Run(":8080")
}
