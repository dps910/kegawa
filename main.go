package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Cannot create window:", err)
	}
	window.SetTitle("title")
	window.SetDefaultSize(800, 600)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create new label widget for gtk window
	l, err := gtk.LabelNew("Hello, I am a label widget in a window")
	if err != nil {
		log.Fatal("Couldn't create label:", err)
	}

	window.Add(l)

	// Add CSS
	cssProv, _ := gtk.CssProviderNew()
	cssProv.LoadFromPath("res/style.css")

	context, _ := l.GetStyleContext()
	context.AddProvider(cssProv, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// Show all widgets contained in gtk window
	window.ShowAll()

	// Start gtk main loop
	gtk.Main()
}
