package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

//Pkg is a "go list -json" result
type Pkg struct {
	Name   string   `json:"Name"`
	Import string   `json:"ImportPath"`
	Deps   []string `json:"Deps"`
}

var mu sync.Mutex
var commonPacks = map[string]int{}

// Searching common dependencies for packages
// example:
// go run main.go encoding/json encoding/xml archive/zip fmt
func main() {

	var wg sync.WaitGroup

	for _, arg := range os.Args[1:] {
		wg.Add(1)
		go getPackageList(arg, &wg)
	}

	wg.Wait()

	fmt.Println("Common dependencies for packages:")
	for pack, count := range commonPacks {
		if count == len(os.Args[1:]) {
			fmt.Println(pack, count)
		}
	}
}

func getPackageList(arg string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("go", []string{"list", "-json", arg}...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Command %s, error - %s", cmd.Args, err)
	}

	var pkg Pkg
	if err := json.Unmarshal(output, &pkg); err != nil {
		panic(err)
	}

	fmt.Println(pkg.Deps)

	mu.Lock()
	for _, pack := range pkg.Deps {
		if _, ok := commonPacks[pack]; !ok {
			commonPacks[pack] = 1
		} else {
			commonPacks[pack]++
		}
	}
	mu.Unlock()
}
