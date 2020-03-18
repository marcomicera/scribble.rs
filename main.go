package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	errorPage *template.Template
	portHTTP  *int
)

func main() {
	// "When you deploy an app through heroku, it does not allow you to specify the port number"
	// https://stackoverflow.com/a/51344239/3927431
	portHTTP = new(int)
	var parseError error
	*portHTTP, parseError = strconv.Atoi(os.Getenv("PORT"))
	if parseError != nil {
		panic(parseError)
	}

	flag.Parse()

	//Setting the seed in order for the petnames to be random.
	rand.Seed(time.Now().UnixNano())

	errorPage, parseError = template.New("").ParseFiles("error.html", "footer.html")
	if parseError != nil {
		panic(parseError)
	}

	setupRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portHTTP), nil))
}

func setupRoutes() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
}
