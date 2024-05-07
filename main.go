package main

import (
	"bytes"
	"fmt"
	goimage "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"golang.org/x/image/bmp"
)

func main() {

	fileBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}

	contentType := http.DetectContentType(fileBytes)

	var img goimage.Image

	switch contentType {
	case "image/png":
		img, err = png.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}

	case "image/gif":
		img, err = gif.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
	}

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, img); err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	bmpImage, err := bmp.Decode(buf)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			r, g, b, _ := bmpImage.At(i, j).RGBA()
			fmt.Printf("#%x%x%x\n", uint8(r), uint8(g), uint8(b))
		}
	}

}
