cities-by-state
===============

Simple service to get cities by state abbreviation using the [Small Business Association's U.S. City & County Web Data API](http://api.sba.gov/doc/geodata.html), [http://api.sba.gov/doc/geodata.html](http://api.sba.gov/doc/geodata.html)

[![GoDoc](https://godoc.org/github.com/savaki/cities-by-state?status.svg)](https://godoc.org/github.com/savaki/cities-by-state)

## Example - Using the API

```
import (
	"github.com/savaki/cities-by-state"
	"json"
	"os"
)

func main() {
	cityService := citiesbystate.New()
	cities, _ := cityService.ByState("California")
	json.NewEncoder(os.Stdout).Encode(cities)
}
```

## Example - Using with a Google Context 

Often requests participate as part of a chain of requests, using Google context simplifies management of that call change.  From the Golang blog:

> At Google, we developed a context package that makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. 


```
import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/cities-by-state"
	"json"
	"os"
	"time"
)

func main() {
	// require this request completes within 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	cities, _ := citiesbystate.WithContext(ctx).ByState("CA"))
	json.NewEncoder(os.Stdout).Encode(cities)
}
```


