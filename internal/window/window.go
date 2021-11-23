package window

import (
	"os"

	"github.com/webview/webview"

	"harmony/internal/extension"
	"harmony/internal/song"
)

// Init creates a new webview window, loads the frontend
// located in the frontend directory, and binds all Go functions
// that are callable from the frontend. It passes es to some of
// the binds that need to use the extensions.
func Init(es []extension.Extension) {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Harmony")
	w.SetSize(800, 600, webview.HintNone)

	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	w.Navigate("file://" + p + "/frontend/index.html")

	w.Bind("go_addSong", song.New)
	w.Bind("go_getSongs", song.GetSongs)
	w.Bind("go_getAudio", getAudio(es))

	w.Run()
}
