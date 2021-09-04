package main

import "github.com/gotk3/gotk3/gtk"

func main() {
	// Init GTK
	gtk.Init(nil)

	// load glade file
	builder, _ := gtk.BuilderNew()
	builder.AddFromFile("higotk.glade")
	// make a window
	obj, _ := builder.GetObject("window1")
	window := obj.(*gtk.Window)
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()

	// get the label
	// with id label_glade
	obj, _ = builder.GetObject("label_glade")
	label_glade := obj.(*gtk.Label)

	// get button object
	// with id button_hi
	obj, _ = builder.GetObject("button_hi")
	// set apropriate object ref
	button_hi := obj.(*gtk.Button)

	// handle its "clicked" event with our simple callback func
	button_hi.Connect("clicked", func() {
		// set label to User Said Hello!
		label_glade.SetText("Hello User!")
	})

	// run GTK Main
	gtk.Main()
}
