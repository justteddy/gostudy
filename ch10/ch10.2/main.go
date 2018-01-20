package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	readArchive(os.Args[1], os.Stdout)
}

func readArchive(filename string, out io.Writer) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".zip":
		r, err := zip.OpenReader(filename)
		if err != nil {
			fmt.Fprintf(out, "zip archive opening error %s\n", err)
		}
		defer r.Close()

		for _, f := range r.File {
			fmt.Fprintf(out, "%s\n", f.Name)
		}
	case ".tar":
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(out, "tar archive opening error %s\n", err)
		}
		defer f.Close()

		r := tar.NewReader(f)

		for {
			header, err := r.Next()
			if err == io.EOF {
				break
			}

			fmt.Fprintf(out, "%s\n", header.Name)
		}
	default:
		fmt.Fprint(out, "Unknown archive format")
	}
}
