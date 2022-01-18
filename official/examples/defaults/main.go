package main

import (
	"context"
	"fmt"
	"log"

	aoe "github.com/theflyingcodr/aoe4-client"
	"github.com/theflyingcodr/aoe4-client/official"
)

// this example uses the default arguments in the client so no filters
// are supplied, these defaults match the website when first opened and show
// the top 100 global players for unranked 1v1 matches.
func main() {
	c := official.NewClient(aoe.NewHTTPClient())

	resp, err := c.NewRequest().Do(context.Background()) // use a proper context in an actual app
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total returned %d\n", resp.Count)
}
