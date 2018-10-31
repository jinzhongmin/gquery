# gquery
Parse and use html files.
For quick and easy use of html, you can manipulate html like jquery.
# example

``` go
package main

import (
	"net/http"
	"github.com/jinzhongmin/gquery"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		doc, _ := gquery.NewDocument("index.html")
		//Handling html files here
		return c.HTML(http.StatusOK, doc.Render())
	})

	e.Start(":80")
}

```
# api
https://gowalker.org/github.com/jinzhongmin/gquery

# postscript

Originated from an idea when I just learned golang, I don't use golang's template engine. Can I generate html content in a way like jquery? So the project was born. The initial prototype is https://github.com/jinzhongmin/elm. Later I found https://github.com/PuerkitoBio/goquery and then found http://github.com/andybalholm/cascadia. So I plan to process the project I just started, as close as possible to the jquery api. When I finished this project initially, I was going to write an example to test it and found that it was not very easy to handle html when using this project. 😂
