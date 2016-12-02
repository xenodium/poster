package writer

import (
	"fmt"
	"html"
	"image"
	"image/color"
	"io"
	"poster/source"
)

func Write(image image.Image, filePaths []string, bgColor string, writer io.Writer) {
	if len(bgColor) == 0 {
		bgColor = "#FFFFFF"
	}

	writer.Write([]byte(fmt.Sprintf(`<body style="font-family:monospace;
                            background-color:%s;
                            font-size:1.1em;
                            font-size:12px;
                            letter-spacing:0.3em;
                            line-height:0.5em;
                            white-space:pre;">`, bgColor)))

	it := source.NewIterator(filePaths)

	for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {
		for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
			// Convert to non-alpha-premultiplied color.
			nrgbaC := color.NRGBAModel.Convert(image.At(x, y)).(color.NRGBA)

			// Restart iterator if we run out of characters.
			if it.Done() {
				it.Reset()
			}

			// Replace newlines, tabs, and non-ascii characters with following character.
			var charPixel = "\n"
			for charPixel = html.EscapeString(*it.Read()); charPixel == "\n" || charPixel == "\t" || !isASCII(charPixel); charPixel = html.EscapeString(*it.Read()) {
				if charPixel == "\n" || charPixel == "\t" {
					charPixel = " "
					break
				}
			}

			writer.Write([]byte(fmt.Sprintf(`<span style="color:rgba(%d, %d, %d, %.2f);">%s</span>`,
				nrgbaC.R,
				nrgbaC.G,
				nrgbaC.B,
				float32(nrgbaC.A)/255.0,
				charPixel)))
		}
		writer.Write([]byte("<br/>\n"))
	}

	writer.Write([]byte(`</body>`))
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}
