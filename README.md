# DuckDuckGo Query Library
[![pipeline status](https://gitlab.com/psheets/ddgquery/badges/master/pipeline.svg)](https://gitlab.com/psheets/ddgquery/commits/master) [![coverage report](https://gitlab.com/psheets/ddgquery/badges/master/coverage.svg)](https://gitlab.com/psheets/ddgquery/commits/master)

ddgquery provides a method to retrieve basic search results from DuckDuckGo. It leverages [goquery][] and its very robust feature set to query DuckDuckGo.

## Installation

    $ go get https://gitlab.com/psheets/ddgquery
    
## Example

```golang
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
```

## License

The [BSD 3-Clause license][bsd]

[goquery]: http://github.com/PuerkitoBio/goquery
[bsd]: http://opensource.org/licenses/BSD-3-Clause