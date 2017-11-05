package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type customSort struct {
	t           []*Track
	currentSort string
}

func (x customSort) Len() int      { return len(x.t) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x customSort) Less(i, j int) bool {
	switch x.currentSort {
	case "Title":
		return x.t[i].Title < x.t[j].Title
	case "Year":
		return x.t[i].Year < x.t[j].Year
	case "Length":
		return x.t[i].Length < x.t[j].Length
	case "Artist":
		return x.t[i].Artist < x.t[j].Artist
	case "Album":
		return x.t[i].Album < x.t[j].Album
	}

	return false
}

var cs = customSort{
	t: []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	},
	currentSort: "Title",
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func sortFunc(w http.ResponseWriter, r *http.Request) {
	sortField := r.FormValue("sort")
	cs.currentSort = sortField
	sort.Sort(cs)

	table, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := table.Execute(w, cs.t); err != nil {
		log.Fatal(err)
	}

}

func main() {

	http.HandleFunc("/", sortFunc)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
