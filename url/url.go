package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, error := url.Parse("https://localhost:80/api/books/12345678?title=sudhir&&name=pola")
	if error != nil {
		log.Fatal(error)
	}
	/* print the parsed contents */
	fmt.Printf("Scheme: %s\n", u.Scheme)
	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Path: %s\n", u.Path)
	q := u.Query()
	fmt.Printf("Length of the map: %v\n", len(q))
	fmt.Printf("Query: %v\n", q)

}
