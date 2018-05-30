package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
  // serve the "dist" folder on root ("/") path
	http.Handle("/", http.FileServer(http.Dir("./dist")))
  // start the appengine server
	appengine.Main()
}
