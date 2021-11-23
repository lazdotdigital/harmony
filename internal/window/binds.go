package window

import (
	"harmony/internal/extension"
)

// getAudio returns a function that takes a URL,
// then passes it to every installed extension's Claim function.
// Upon an extension returning true, it called that function's
// GetAudio function and returns the temp audio file path.
func getAudio(es []extension.Extension) func(string) string {
	return func(url string) string {
		for _, e := range es {
			if e.Exports.Claim(url) {
				return e.Exports.GetAudio(url)
			}
		}

		return ""
	}
}
