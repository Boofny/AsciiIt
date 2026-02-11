package main

import (
	// "bytes"
	// "fmt"
	// "image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
)

const URL = "https://picsum.photos/200/300"

func main() {
	chans()
}

func norm(){
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	os.WriteFile("output.jpg", b, 0664)
}

func chans(){
	ch := make(chan *http.Response)
	go func (url string)  {
		resp, err := http.Get(URL)
		if err != nil {
			panic(err)
		}
		ch <- resp
	}(URL)

	r := <-ch
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	os.WriteFile("output.jpg", b, 0664)
}
