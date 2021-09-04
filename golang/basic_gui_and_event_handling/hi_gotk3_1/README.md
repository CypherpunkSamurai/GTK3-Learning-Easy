# Hi GoTK3 - 1
In this tutorial we try to compile a simple basic window with a label in the center.

![screenshot](https://user-images.githubusercontent.com/66906402/132097084-53abac71-7b0d-4eac-b7b7-d7da6316a8d9.png)


## How to compile

* Open this folder
* run `go build`
* It should produce a binary, run it!

## Code

We import gotk3 in gotk3_1.go#6 (line 6) with `"github.com/gotk3/gotk3/gtk"`

Then we make a main function for out go app.

On line 11 we run `gtk.Init(nil)` to init the GTK library

On line 13 we run `win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)` to make a new window object with `WINDOW_TOPLEVEL` attribute. We also handle error with `err`

On line 17 we run `win.SetTitle("Hi GoTK3")` to set the window title to "Hi GoTK3".

On line 18 we run `win.Connect("destroy", func() {...}` to connect the destroy event of the window to the function (called callback function). Whenever the window is closed this function will be called.

On line 19 we run `gtk.MainQuit()` to quit the GTK main process.

On line 22 we run `l, err := gtk.LabelNew("Hi GoTK3!")` to make a new label with text "Hi GoTK3!". We handle errors again.

On line 27 we run `win.Add(l)` we add label to the window.

On line 29 we run `win.SetDefaultSize(800, 600)` to set the window default size to height=800 and width=600.

On line 31 we run `win.ShowAll()` to make the window show.

On line 33 we run `gtk.Main()` to run the gtk main process.

## What are Event Handlers

Having come from C# background since i was 12, event handlers seemed easy. But for some it might seem confusing.

Event handlers are method functions (normally ones that dont return data) that run when events occur. This should used by the programmer to process stuff.

Events like clicking a button, scrolling a scroll bar, resizing the window, right clicking on the form, hovering over the form, or keyboard keypress.

## Notes

To better understand try editing the label and window titles. Try making this later without looking throught the code.
