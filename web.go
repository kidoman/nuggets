package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

var pike image.Image

func init() {
	f, err := os.Open("pike.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pike, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/love.png", love)
	http.HandleFunc("/", yum)
}

func yum(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(w, f)
}

func love(w http.ResponseWriter, r *http.Request) {
	png.Encode(w, pike)
}
