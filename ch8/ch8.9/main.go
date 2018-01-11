package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type dirInfo struct {
	name string
	size int64
}

//!+
func main() {
	flag.Parse()
	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan *dirInfo)
	directories := make(map[string]*dirInfo)
	var wg sync.WaitGroup

	for _, root := range roots {
		wg.Add(1)
		directories[root] = &dirInfo{root, 0}
		go walkDir(root, root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	for tempDirInfo := range fileSizes {
		directories[tempDirInfo.name].size += tempDirInfo.size
	}
	printDiskUsage(roots, directories) //final totals
}

func printDiskUsage(roots []string, directories map[string]*dirInfo) {
	for _, root := range roots {
		fmt.Printf("%s\t%.1f GB\n", directories[root].name, float64(directories[root].size)/1e9)
	}
}

func walkDir(dir string, root string, wg *sync.WaitGroup, fileSizes chan<- *dirInfo) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, root, wg, fileSizes)
		} else {
			fileSizes <- &dirInfo{root, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
