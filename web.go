package main

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nfnt/resize"

	"github.com/kidoman/nuggets/types"
	_ "github.com/kidoman/nuggets/types/set1"
	// _ "github.com/kidoman/nuggets/types/set2"
)

const (
	width  = 100
	height = 100
	startX = (624 - width - 20)
	startY = 25
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
	target := image.NewRGBA(image.Rect(0, 0, 624, 480))
	draw.Draw(target, pike.Bounds(), pike, image.ZP, draw.Src)

	start := time.Now()

	for i, url := range types.Nuggets {
		x := startX - (i/4)*120
		y := startY + (i%4)*110

		nugget, err := fetchReziedImage(url)
		if err != nil {
			panic(err)
		}
		draw.Draw(target, image.Rect(x, y, x+width, y+height), nugget, image.ZP, draw.Src)
	}

	log.Printf("total time taken: %v", time.Since(start))

	png.Encode(w, target)
}

func fetchReziedImage(url string) (image.Image, error) {
	log.Printf("fetching %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	nugget, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("fetched %v", url)
	nugget = resize.Resize(width, height, nugget, resize.Lanczos3)
	return nugget, nil
}
