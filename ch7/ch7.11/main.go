package main

import "net/http"
import "log"
import "fmt"
import "strconv"

type database map[string]dollars

func (d *database) delete(name string) bool {
	db := *d
	if _, ok := db[name]; !ok {
		return false
	}

	delete(db, name)
	return true
}

func (d *database) add(name string, price float32) bool {
	db := *d
	db[name] = dollars(price)

	return true
}

func (d database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := d[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 60, "socks": 10}
	mux := http.NewServeMux()

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	mux.HandleFunc("/add", func(w http.ResponseWriter, req *http.Request) {
		item := req.URL.Query().Get("item")
		price := req.URL.Query().Get("price")
		if price == "" || item == "" {
			fmt.Fprintln(w, "Attribute error")
			return
		}

		cost, err := strconv.ParseFloat(price, 32)
		if err != nil {
			fmt.Fprintf(w, "Price error - %s", err)
			return
		}

		db.add(item, float32(cost))
		fmt.Fprintf(w, "Item %s has added with price %s", item, dollars(cost))
	})
	mux.HandleFunc("/delete", func(w http.ResponseWriter, req *http.Request) {
		item := req.URL.Query().Get("item")
		if item == "" {
			fmt.Fprintln(w, "Attribute error")
			return
		}

		if !db.delete(item) {
			fmt.Fprintf(w, "Can not delete item - %s", item)
			return
		}

		fmt.Fprintf(w, "Item - %s was deleted", item)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
