package main

import "github.com/gotk3/gotk3/gtk"

func main() {
	// Init GTK
	gtk.Init(nil)

	// Make new GTK builder
	builder, _ := gtk.BuilderNew()

	// User builder to load from glade file
	builder.AddFromFile("higotk.glade")

	// get a window object
	obj, _ := builder.GetObject("window1")

	// make new window
	window := obj.(*gtk.Window)

	// connect window destroy event to gtk MainQuit method
	window.Connect("destroy", gtk.MainQuit)

	// windows size is set from the glade

	// display window
	window.ShowAll()

	// run GTK Main
	gtk.Main()
}
