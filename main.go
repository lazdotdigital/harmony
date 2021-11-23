package main

import (
	_ "embed"
	"harmony/internal/extension"
	"harmony/internal/window"
)

// main loads all extensions and calls window.Init
func main() {
	e, err := extension.New("./builtin-extensions/bandcamp.js")
	if err != nil {
		panic(err)
	}

	e2, err := extension.New("./builtin-extensions/youtube.js")
	if err != nil {
		panic(err)
	}

	es := []extension.Extension{e, e2}
	window.Init(es)
}
