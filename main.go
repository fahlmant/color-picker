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
		return
	}

	imgConfig, _, err := goimage.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		fmt.Printf("Img Config Error: %+v\n", err)
		return
	}

	for i := 0; i < imgConfig.Bounds().Dx(); i++ {
		for j := 0; j < imgConfig.Bounds().Dy(); j++ {
			r, g, b, _ := bmpImage.At(i, j).RGBA()
			// RGBA returns a number between 0 and 65535. 255 * 257 = 65535, so we divide by 257 to get a number between 0 and 255
			fmt.Printf("#%02x%02x%02x\n", uint32((float64(r) / 257)), uint32((float64(g) / 257)), uint32((float64(b) / 257)))
		}
	}

}
