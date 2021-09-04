package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

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
	obj, _ = builder.GetObject("label_glade")
	label_glade := obj.(*gtk.Label)

	// get text box
	obj, _ = builder.GetObject("entry_text1")
	entry_text1 := obj.(*gtk.Entry)

	// get button
	obj, _ = builder.GetObject("button_hi")
	button_hi := obj.(*gtk.Button)
	button_hi.Connect("clicked", func() {
		// set label to User Said Hello!
		entered_text, err := entry_text1.GetText()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		label_glade.SetText(entered_text)
	})

	// run GTK Main
	gtk.Main()
}
