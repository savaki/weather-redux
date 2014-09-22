openweathermap 
==============

Simple Golang client into the [openweathermap](http://openweathermap.org/api) API

[![GoDoc](https://godoc.org/github.com/savaki/openweathermap?status.svg)](https://godoc.org/github.com/savaki/openweathermap)

## Example - Anonymous Usage

```
import "github.com/savaki/openweathermap"

func main() {
	weatherService := openweathermap.New()
	forecast, err := weatherService.FindByCity("San Francisco")
}
```

## Example - Using an API key

```
import (
	"github.com/savaki/openweathermap"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	weatherService := openweathermap.WithApiKey(apiKey))
	forecast, err := weatherService.FindByCity("San Francisco")
}
```

## Example - Using with a Google Context and an (optional) API key

Often requests participate as part of a chain of requests, using Google context simplifies management of that call change.  From the Golang blog:

> At Google, we developed a context package that makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. 


```
import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/openweathermap"
	"os"
	"time"
)

func main() {
	// require this request completes within 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	apiKey := os.Getenv("API_KEY")
	
	weatherService := openweathermap.WithContext(ctx, apiKey)
	forecast, err := weatherService.FindByCity("San Francisco")
}
```


