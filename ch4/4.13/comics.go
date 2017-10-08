package main

import (
	"html/template"
	"log"
	"net/http"

	"gopl.io/ch4/xkcd"
)

func main() {
	http.HandleFunc("/", showComics)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func showComics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	comicsID := r.URL.Query().Get("id")
	result, err := xkcd.SearchComics(comicsID)
	if err != nil {
		log.Fatal(err)
	}

	var comicsPage = template.Must(template.New("comicsPage").
		Parse(xkcd.HtmlTemplate))
	if err := comicsPage.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
