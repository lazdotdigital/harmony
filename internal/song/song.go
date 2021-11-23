package song

import (
	"encoding/json"
	"io/ioutil"
)

// Song defines a saved song that is stored by Harmony.
type Song struct {
	// The title of the song
	Title string `json:"title"`
	// Who the song is by
	By string `json:"by"`
	// The path to the song.
	AudioPath string `json:"audioPath"`
}

// New takes title, by, and audioTmp, then permanently
// saves it to the songs directory.
func New(title, by, audioTmp string) (s Song, err error) {
	s.Title = title
	s.By = by

	s.AudioPath, err = save(audioTmp)
	if err != nil {
		return
	}

	ss, err := GetSongs()
	if err != nil {
		return
	}

	ss = append(ss, s)
	data, err := json.Marshal(ss)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("./songs/songs.json", data, 0)

	return
}

// GetSongs returns the parsed songs JSON file.
func GetSongs() (ss []Song, err error) {
	songsJson, err := ioutil.ReadFile("./songs/songs.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(songsJson, &ss)
	return
}
