package ddgcrawl
import (
	"log"
	"fmt"
	"strings"
	"net/url"
	
	"github.com/PuerkitoBio/goquery"

)

type Qresult struct {
	Title 	string
	Info 	string
	Ref		string
}

func Query(q string) ([]Qresult, string, int) {

	qf := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", url.QueryEscape(q))
	doc, err := goquery.NewDocument(qf)

	results:= []Qresult{}
	if err != nil {
	    log.Fatal(err)
	}

	doc.Find(".web-result").Each(func(i int, s *goquery.Selection) {
	    // For each item found, get the band and title
	if (i < 2)
		titleNode:=s.Find(".result__a")
		info:=s.Find(".result__snippet").Text()
		title := titleNode.Text()

	    ref, _ := url.QueryUnescape(strings.TrimPrefix(titleNode.Nodes[0].Attr[2].Val, "/l/?kh=-1&uddg="))
	    results = append(results[:], Qresult{title, info, ref})

	  })

	return results, qf
}

 