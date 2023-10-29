package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO:
//  r: Download the file
//  w: Progress counter: Measure in real time the number of mb downloaded
//  w: Write the file into our local filesystem
//  w: Write the file into our archive

type Counter struct {
  total uint64
}

func (c *Counter) Write(b []byte) (int, error) {
  c.total += uint64(len(b))
  progress := float64(c.total) / (1024 * 1024)
  fmt.Printf("\rDownloading %f MB...", progress)
  return len(b), nil
}

func main() {
  var URL string = `http://storage.googleapis.com/books/ngrams/books/googlebooks-eng-all-5gram-20120701-0.gz` 
  fmt.Println("Fetching: ", URL)
  res, err := http.Get(URL)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()

  // Checking if the file has a valid header
  if res.Header.Get("Content-Encoding") != "gzip" {
    panic("The file is no a valid gip.")
  }

  // Downloading the file in the local system
  localF, err := os.OpenFile("download-5gram.txt", os.O_CREATE|os.O_WRONLY, 0600)
  if err != nil {
    panic(err)
  }
  defer localF.Close()

  // Decompress...
  dec, err := gzip.NewReader(res.Body)
  if err != nil {
    panic(err)
  }

  // Copy res.Body into local
  if _, err := io.Copy(localF, io.TeeReader(dec, &Counter{})); err != nil {
    panic(err)
  }
}
