package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo

	text       string
	UppperCase string
}

func ConvertToUpper(s string) string {
	return strings.ToUpper(s)
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().Body(app.Script().Src("https://cdn.tailwindcss.com"),
		app.H1().Body(
			app.Text("Input Text"),
		).Class("text-center text-white text-4xl font-bold"),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.text).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.UppperCase)).Class("w-full p-2 border rounded bg-gray-400"),
		).Style("margin-top", "1em").Style("background-color", "black").Style("color", "white"),
		app.P().Body(
			app.Text("TO UPPERCASE click enter: "),
		).Class("text-center text-white text-md font-bold"),
		app.P().Body(
			app.Text(ConvertToUpper(h.UppperCase)),
		).Class("text-center text-white text-md font-bold"),
	).Class("flex flex-col flex-wrap items-center gap-8 justify-center h-screen w-screen bg-gradient-to-r from-gray-900 to-gray-800")
}
func main() {
	// Components routing:
	app.Route("/", &hello{})
	app.Route("/hello", &hello{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
