package xkcd

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const XkcdURL = "https://xkcd.com/%s/info.0.json"
const HtmlTemplate = `
<center>
<h1>{{.Title}}</h1><br>
<p>{{.PostDate}}</p><br>
<img src="{{.Img}}"><br>
<a href="{{.Next}}"><h4>Next!</h4></a><br>
</center>
`

type Comics struct {
	Num        int
	Month      string
	Year       string
	Day        string
	Title      string
	Alt        string
	Img        string
	Transcript string
}

// PostDate return comics post date
func (c *Comics) PostDate() string {
	month, _ := strconv.Atoi(c.Month)
	t, err := time.Parse("January 2, 2006", fmt.Sprintf("%s %s, %s", time.Month(month), c.Day, c.Year))
	if err != nil {
		log.Fatal(err)
	}

	return t.Format("02 Jan 06")
}

// Next return comics post date
func (c *Comics) Next() string {
	next := c.Num + 1

	return fmt.Sprintf("/?id=%d", next)
}
