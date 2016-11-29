package main

import (
	"fmt"
	"html"
	"image/color"
	"image/png"
	"nfl/source"
	"os"
)

func main() {
	infile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		panic(err)
	}

	fmt.Println(`<body style="font-family:monospace;
                            font-size:1.1em;
                            font-size:12px;
                            letter-spacing:0.3em;
                            line-height:0.5em;
                            white-space:pre;">`)

	it := source.NewIterator(source.ListFiles("/Users/tuco/stuff/active/code/gopath/src/nfl/"))

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// Convert to non-alpha-premultiplied color.
			nrgbaC := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)

			// Restart iterator if we run out of characters.
			if it.Done() {
				it.Reset()
			}

			var charPixel = "\n"
			// Replace newlines, tabs, and non-ascii characters with following character.
			for charPixel = html.EscapeString(*it.Read()); charPixel == "\n" || charPixel == "\t" || !isASCII(charPixel); charPixel = html.EscapeString(*it.Read()) {
				if charPixel == "\n" || charPixel == "\t" {
					charPixel = " "
					break
				}
			}
			fmt.Printf(`<span style="color:rgba(%d, %d, %d, %.2f);">%s</span>`, nrgbaC.R, nrgbaC.G, nrgbaC.B, float32(nrgbaC.A)/255.0, charPixel)
		}
		fmt.Printf("<br/>\n")
	}
	fmt.Println(`</body>`)
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}
