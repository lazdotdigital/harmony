package extension

import (
	"io/ioutil"
	"net/http"
)

// fetch is a helper function that returns the body of url.
func fetch(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	data, err = ioutil.ReadAll(res.Body)
	return
}
