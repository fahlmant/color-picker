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
	"sort"
)

func main() {

	// Flag for specifying how many results to return. For example, --results 5 shows the top 5 hex values by pixel count
	numResults := flag.Int("results", 10, "Specify the number of results to return")
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

	// Create a map for tracking hexcode frequencies
	hexCodeFreq := make(map[string]int)

	// Find the hexcode for each pixel and increment the frequency in the map for that hex code
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			// RGBA returns a number between 0 and 65535. 255 * 257 = 65535, so we divide by 257 to get a number between 0 and 255
			hexCode := fmt.Sprintf("#%02x%02x%02x\n", uint32((float64(r) / 257)), uint32((float64(g) / 257)), uint32((float64(b) / 257)))
			// Increases the counter for a hex string for each occurence
			hexCodeFreq[hexCode] += 1
		}
	}

	// Create a Sorted Hex Code Frequncy array and fill it with all hex codes by using the keys from the frequency map
	sortedHexCodeFreq := make([]string, 0, len(hexCodeFreq))
	for k := range hexCodeFreq {
		sortedHexCodeFreq = append(sortedHexCodeFreq, k)
	}

	// Sort the frequency array by using the values from the frequency map
	sort.Slice(sortedHexCodeFreq, func(i, j int) bool {
		return hexCodeFreq[sortedHexCodeFreq[i]] < hexCodeFreq[sortedHexCodeFreq[j]]
	})

	firstElementIndex := len(sortedHexCodeFreq) - *numResults
	if firstElementIndex >= len(sortedHexCodeFreq) {
		firstElementIndex = len(sortedHexCodeFreq)
	}

	// Print the top  N hex values used in the image
	for _, v := range sortedHexCodeFreq[firstElementIndex:] {
		fmt.Printf("%s", v)
	}

}
