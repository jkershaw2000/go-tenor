# go-tenor

A tenor.com API wrapper for the Go programming language.

## Installation
Configure $GOPATH correctly then


`go get github.com/jkershaw2000/go-tenor`
## Usage

### Search for a GIF
Allow the user to search for a gif hosted on tenor.
`GetSearch` takes 4 parameters

`query (string)` - the search term

`filter ("off", "low", "medium", "high)` - content filtering

`locale (2 letter ISO country code)` - what langauge the search is in

`limit (int, -1 for no limit )` - max number of search results  

```go 
package main

import (
	"fmt"
	"github.com/jkershaw2000/go-tenor"
)
const (
    APIKey = "YOUR_API_KEY"
)

func main() {
	t := gotenor.NewTenor(APIKey)

	data, err := t.GetSearch("Query", "", "", -1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

```

### Get single ID
Returns all data about a GIF.
`GetByID` takes a single parameter

 `id (string)` - gif id
```go
package main

import (
    "fmt"
    "github.com/jkershaw2000/go-tenor"
)
const (
    APIKey = "YOUR_API_KEY"
)

func main() {
    tenor := gotenor.NewTenor(APIKey)

    gifID := "8776030"
    data, err := tenor.GetByID(gifID)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    fmt.Println(data)
}
```

### Get gif URL from tenorData structure
There are two helper methods to get just the URLs for gifs to share.

`GetGifURL(data tenorData)` - returns the first 'gif' url

`GetAllGifURLS(data tenorData)` - returns all 'gif' urls found

```go
package main

import (
    "fmt"
    "github.com/jkershaw2000/go-tenor"
)
const (
    APIKey = "YOUR_API_KEY"
)

func main() {
    tenor := gotenor.NewTenor(APIKey)
    gifID := "8776030,17437428"
    data, err := tenor.GetByID(gifID)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    fmt.Println(libtenor.GetGifURL(*data)) // single URL
    fmt.Println(libtenor.GetAllGifURLS(*data)) // All urls
}

```

## Credits
A lot of the code was adapted from sanzaru's go-giphy which can be found [here](https://github.com/sanzaru/go-giphy).
