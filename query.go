// Package ddgquery provides a method to retrieve basic search results from DuckDuckGo
// ddgquery uses goquery from PuerkitoBio

package ddgquery
import (
	"log"
	"fmt"
	"strings"
	"net/url"
	
	"github.com/PuerkitoBio/goquery"

)
// A Qresult holds the returned query data
type Qresult struct {
	Title 	string
	Info 	string
	Ref		string
}

// Requests the query and puts the results into an array
func Query(q string, it int) ([]Qresult, string,) {

	qf := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", url.QueryEscape(q))
	doc, err := goquery.NewDocument(qf)

	results:= []Qresult{}

	if err != nil {
	    log.Fatal(err)
	}

	sel := doc.Find(".web-result")

	  for i := range sel.Nodes {
	  	// Break loop once required amount of results are add
		if ( it == len(results) ) { break }

		single := sel.Eq(i)
		titleNode:=single.Find(".result__a")
		info:=single.Find(".result__snippet").Text()
		title := titleNode.Text()
	    ref, _ := url.QueryUnescape(strings.TrimPrefix(titleNode.Nodes[0].Attr[2].Val, "/l/?kh=-1&uddg="))

	    results = append(results[:], Qresult{title, info, ref})

	  }

	// Return array of results and formated query used to get the results
	return results, qf
}

 