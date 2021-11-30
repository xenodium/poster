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
	bgColor    string
}

func parseArgs() args {
	args := args{}
	flag.StringVar(&args.searchPath, "searchpath", "", "`path` to search for source files recursively")
	flag.StringVar(&args.bgColor, "bgcolor", "", "background `\"color\"`, for example \"#FFFFFF\" (optional)")

	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n  image.png\n\n")
	}

	flag.Parse()

	args.imagePath = flag.Arg(0)

	if args.searchPath == "" || args.imagePath == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	return args
}

func main() {
	args := parseArgs()

	imageFile, err := os.Open(args.imagePath)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer imageFile.Close()

	image, err := png.Decode(imageFile)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	outputFileName := outputFileName(args.imagePath)
	f, err := os.Create(outputFileName)
	writer.Write(image, source.ListFiles(args.searchPath), args.bgColor, f)
	defer f.Close()

	fmt.Printf("Generated %s\n", outputFileName)
}

func outputFileName(inputFileName string) string {
	return fmt.Sprintf("%s.%s", inputFileName[0:len(inputFileName)-len(filepath.Ext(inputFileName))], "html")
}
