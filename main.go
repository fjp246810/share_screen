package main

import (
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/saljam/mjpeg"
	"github.com/vova616/screenshot"
)

func main() {
	// create the mjpeg stream
	stream := mjpeg.NewStream()

	// start capturing
	go func(stream *mjpeg.Stream) {
		for {
			img, err := screenshot.CaptureScreen()
			if err != nil {
				panic(err)
			}
			file, err := os.Create("./tmp.jpeg")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			jpeg.Encode(file, img, nil)
			buf, err := ioutil.ReadFile("tmp.jpeg")
			if err != nil {
				panic(err)
			}
			stream.UpdateJPEG(buf)
		}
	}(stream)

	http.HandleFunc("/mjpeg", stream.ServeHTTP)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := "<img width='100%' height='100%' src='/mjpeg'>"
		w.Header().Add("Content-Type", "text/html;charset=utf-8")
		w.Write([]byte(content))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
