package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}

func (p *Point) add(offset float64) {
	p.x += offset
	p.y += offset
}

func main() {
	p := Point{3, 3}

	use(p)
}

func use(thing interface{}) {
	original, ok := thing.(Point)

	if !ok {
		println("FUUUUCK!")
	}
	original.add(3)

	fmt.Println(original)
}
