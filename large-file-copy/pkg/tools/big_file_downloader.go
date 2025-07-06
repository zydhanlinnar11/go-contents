package tools

import (
	"io"
	"log"
	"net/http"
)

const (
	Url = "https://link.testfile.org/500MB"
)

func DownloadBigFile() (io.ReadCloser, error) {
	req, err := http.Get(Url)
	if err != nil {
		log.Fatal(err)
	}

	return req.Body, nil
}
