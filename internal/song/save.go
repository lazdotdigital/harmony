package song

import (
	"io"
	"os"

	"github.com/teris-io/shortid"
)

// save takes the audio saved at audioTmp,
// copies it into the songs directory with a
// randomly generated name, then deletes the temp file.
// It returns the path to the newly saved song.
func save(audioTmp string) (path string, err error) {
	defer os.Remove(audioTmp)

	id, err := shortid.Generate()
	if err != nil {
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		return
	}

	path = wd + "/songs/" + id + ".mp3"
	destination, err := os.Create(path)
	if err != nil {
		return
	}
	defer destination.Close()

	source, err := os.Open(audioTmp)
	if err != nil {
		return
	}

	_, err = io.Copy(destination, source)

	return
}
