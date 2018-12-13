package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

// PrintDownloadPercent is to show progress
func PrintDownloadPercent(done chan int64, path string, total int64) {

	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100

			fmt.Printf("%.0f", percent)
			fmt.Print("% ... ")
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

// HTTPDownload is to get file from HTTP
func HTTPDownload(url string, dest string) (err error) {

	file := path.Base(url)

	var path bytes.Buffer
	path.WriteString(dest)

	start := time.Now()

	out, err := os.Create(path.String())

	if err != nil {
		fmt.Println(path.String())
		return
	}

	defer out.Close()

	headResp, err := http.Head(url)

	if err != nil {
		return
	}

	defer headResp.Body.Close()

	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if err != nil {
		return
	}

	done := make(chan int64)

	go PrintDownloadPercent(done, path.String(), int64(size))

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)

	if err != nil {
		return
	}

	done <- n

	elapsed := time.Since(start)
	fmt.Println("completed!", file, elapsed)
	return
}
