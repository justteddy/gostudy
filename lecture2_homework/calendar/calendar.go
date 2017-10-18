package main

import (
	"fmt"
	"time"
)

func main() {
	parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", "02"))
	calendar := NewCalendar(parsed)
	calendar.CurrentQuarter()
}

// NewCalendar time type
type NewCalendar interface{}

// CurrentQuarter show quarter
func (n NewCalendar) CurrentQuarter() {
	fmt.Println(123)
}
