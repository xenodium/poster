package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"poster/source"
	"poster/writer"
)

type args struct {
	imagePath  string
	searchPath string
}

func parseArgs() args {
	args := args{}
	flag.StringVar(&args.imagePath, "imagepath", "", "path to image file")
	flag.StringVar(&args.searchPath, "searchpath", "", "path to search for source files (recursively)")
	flag.Parse()
	if len(args.imagePath) == 0 || len(args.searchPath) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	return args
}

func main() {
	args := parseArgs()

	imageFile, err := os.Open(args.imagePath)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()

	// Support png only for now.
	image, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}

	outputFileName := outputFileName(args.imagePath)
	f, err := os.Create(outputFileName)
	writer.Write(image, source.ListFiles(args.searchPath), f)
	defer f.Close()

	fmt.Printf("Generated %s\n", outputFileName)
}

func outputFileName(inputFileName string) string {
	return fmt.Sprintf("%s.%s", inputFileName[0:len(inputFileName)-len(filepath.Ext(inputFileName))], "html")
}
