package extension

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kkdai/youtube/v2"
	"github.com/robertkrimen/otto"
)

// Exports defines the attributes and functions that
// can be set by an extension.
type Exports struct {
	// The name of the extension.
	Name string
	// Claim is called when a new song is added.
	// It takes the song's URL, and returns true
	// if the extension should handle the URL.
	Claim func(string) bool
	// GetAudio is called if Claim returned true.
	// It takes the song's URL, and returns a path
	// to a temp file generated containing the song.
	GetAudio func(string) string
}

// setApi binds the Go functions callable from the VM.
// exports, print, fetchString, fetchBytes, youtubeToMP3
func setApi(vm *otto.Otto, exports *Exports) {
	vm.Set("exports", exports)
	vm.Set("print", print(exports))
	vm.Set("fetchString", fetchString)
	vm.Set("fetchBytes", fetchBytes)
	vm.Set("youtubeToMP3", youtubeToMP3)
}

// youtubeToMP3 takes a YouTube video ID, and returns
// the path to a temp mp3 file containing the video.
func youtubeToMP3(id string) string {
	client := youtube.Client{}
	video, err := client.GetVideo(id)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(file, stream); err != nil {
		panic(err)
	}

	return file.Name()
}

// print returns a function that allows extensions to
// log to the console. Mainly useful for debugging extensions.
func print(e *Exports) func(string) {
	return func(text string) {
		fmt.Printf("EXTENSION %v: %v\n", e.Name, text)
	}
}

// fetchBytes downloads the content at url,
// saves the content of it at a temp file,
// and returns the path to that temp file.
// NOTE: the temp file must be manually removed after it's used.
func fetchBytes(url string) string {
	data, err := fetch(url)
	if err != nil {
		panic(err)
	}

	tmp, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		panic(err)
	}

	if _, err = tmp.Write(data); err != nil {
		panic(err)
	}

	return tmp.Name()
}

// fetchString returns the contents of url as a string.
func fetchString(url string) string {
	data, err := fetch(url)
	if err != nil {
		panic(err)
	}

	return string(data)
}
