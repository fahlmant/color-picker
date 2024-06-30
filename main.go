package main

import (
	"bytes"
	"flag"
	"fmt"
	goimage "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	gq "github.com/cascax/colorthief-go"
)

func main() {

	// Flag for specifying how many results to return. For example, --results 5 shows the top 5 hex values by pixel count
	numResults := flag.Int("results", 16, "Specify the number of results to return")
	flag.Parse()

	// Get file path
	fileName := flag.Args()
	if len(fileName) > 1 || len(fileName) == 0 {
		fmt.Println("Expected one file, exiting")
		return
	}

	// Read the file contents into a byte array
	fileBytes, err := os.ReadFile(fileName[0])
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}

	// Attempt to extract the content type of the file, looking for an image format
	contentType := http.DetectContentType(fileBytes)

	var img goimage.Image

	// Decode PNGs, JPEGs and GIFs into Go Image type
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

	p, err := gq.GetPalette(img, *numResults)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, color := range p {
		r, g, b, _ := color.RGBA()
		fmt.Printf("#%02x%02x%02x\n", uint32((float64(r) / 257)), uint32((float64(g) / 257)), uint32((float64(b) / 257)))
	}

}
