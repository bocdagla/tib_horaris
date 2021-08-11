package downlader

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string, chanel chan<- error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		chanel <- err
		return
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		chanel <- err
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		chanel <- err
		return
	}
}
