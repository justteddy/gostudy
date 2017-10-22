package main

// Placable interface for game places
type Placable interface {
	look() string
	take(string) string
}
