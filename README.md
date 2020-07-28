# go-tenor

A tenor.com API wrapper for the Go programming language.

## Usage

### Get single ID
Returns all data about a GIF.
`GetByID` takes a single parameter, `id (string)` which is the gif ID.
```go
package main


import (
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
        fmt.Println("Error: ". err)
    }
    fmt.Println(data)
}
```
