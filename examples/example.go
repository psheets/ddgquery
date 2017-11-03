package main

import (
	"fmt"
	"gitlab.com/psheets/ddgquery"
)

func main() {

	// Provide query the search term and how many results you want
	// query returns an array of results and the query URL
	results, query := ddgquery.Query("test", 3)

	fmt.Println(query)

	for _, r := range results {
		fmt.Printf("\nTitle: %s \nInfo: %s \nRef: %s \n\n", r.Title, r.Info, r.Ref)
	}
}
	