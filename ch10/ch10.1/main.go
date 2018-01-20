package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var outFmt string
	var filename string
	flag.StringVar(&outFmt, "format", "jpeg", "output format")
	flag.StringVar(&filename, "filename", "test.png", "file to convert")

	flag.Parse()

	//input file
	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "input file open error - %s", err)
	}
	defer inputFile.Close()

	//output file
	extension := filepath.Ext(filename)
	resultFile, err := os.Create("result/" + filename[0:len(filename)-len(extension)] + "." + outFmt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "result file open error - %s", err)
	}
	defer resultFile.Close()

	if err := convertImage(bufio.NewReader(inputFile), bufio.NewWriter(resultFile), outFmt); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func convertImage(in io.Reader, out io.Writer, outFmt string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "Input format =", kind)

	switch outFmt {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return errors.New("unknown file format")
	}
}
