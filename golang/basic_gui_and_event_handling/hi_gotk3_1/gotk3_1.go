package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// init Gtk
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Error:", err)
	}
	win.SetTitle("Hi GoTK3")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("Hi GoTK3!")
	if err != nil {
		log.Fatal("Error:", err)
	}

	win.Add(l)

	win.SetDefaultSize(800, 600)

	win.ShowAll()

	gtk.Main()
}
