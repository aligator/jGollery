package main

import . "jGollery/controller"

func main() {
	webController := NewWebController("static", "gallery")
	webController.Run(":8080")
}
