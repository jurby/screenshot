package main

import (
	"io/ioutil"

	"github.com/pressly/screenshot"
)

/*
 * Example client usage
 * TODO: make more complex examples
 */

func main() {
	client := screenshot.New("http://localhost:3000")
	window := &screenshot.Window{
		Width:  1000,
		Height: 1000,
	}

	png, err := client.PNG("https://golang.org/pkg/fmt/", screenshot.ScreenshotOptions{
		Window: window,
	})

	defer png.Close()

	if err != nil {
		panic(err)
	}

	pngB, err := ioutil.ReadAll(png)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./img.png", pngB, 0644)

	pdf, err := client.PDF("https://golang.org/pkg/fmt/", screenshot.PdfOptions{
		PrintBackground: true,
	})

	defer pdf.Close()

	if err != nil {
		panic(err)
	}

	pdfB, err := ioutil.ReadAll(pdf)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./doc.pdf", pdfB, 0644)

	jpeg, err := client.JPEG("https://golang.org/pkg/fmt/", screenshot.ScreenshotOptions{}, 100)

	defer jpeg.Close()

	if err != nil {
		panic(err)
	}

	jpegB, err := ioutil.ReadAll(jpeg)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./img.jpeg", jpegB, 0644)

}