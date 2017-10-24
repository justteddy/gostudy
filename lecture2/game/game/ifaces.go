package main

// Placable interface for game places
type Placable interface {
	look() string
	oncome() string
	put(string) string
}
