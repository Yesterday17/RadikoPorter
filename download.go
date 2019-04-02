package radiko

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	retry     = 5
	coroutine = 64
)

var ch = make(chan struct{}, coroutine)

// Download - download m3u8 to specified directory
func (t M3U8ChunkList) Download(path string) []string {
	var wg sync.WaitGroup
	errChunks := []string{}

	for _, chunk := range t {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()

			var err error
			for i := 0; i < retry; i++ {
				ch <- struct{}{}
				err = download(link, path)
				<-ch

				if err != nil {
					break
				}
			}

			if err != nil {
				// Failed to download this chunk
				errChunks = append(errChunks, link)
			}
		}(chunk)
	}

	wg.Wait()
	return errChunks
}

func download(link, path string) error {
	client := &http.Client{}
	res, err := client.Get(link)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, fileName := filepath.Split(link)
	file, err := os.Create(filepath.Join(path, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	return err
}
