// Package extension provides a clean wrapper around the Otto JS engine.
// It provides an API that can be called from within extensions,
// as well as defining the extension's exports.
package extension

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

// Extension is a JavaScript VM and the exports defined within the extension.
type Extension struct {
	vm      *otto.Otto
	Exports *Exports
}

// New creates a new Extension by opening the JS file located at path,
// and executing it inside of the vm.
func New(path string) (e Extension, err error) {
	if err != nil {
		panic(err)
	}

	e.vm = otto.New()
	e.Exports = &Exports{}

	setApi(e.vm, e.Exports)

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if _, err = e.vm.Run(src); err != nil {
		return
	}

	return
}
