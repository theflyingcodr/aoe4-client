package main

import (
	"context"
	"fmt"
	"log"

	aoe "github.com/theflyingcodr/aoe4-client"
	"github.com/theflyingcodr/aoe4-client/official"
)

// The api will throttle badly when issuing a lot of requests, you will notice
// this will hit 10 quickly before slowing to a crawl.
func main() {
	c := official.NewClient(aoe.NewHTTPClient())
	totalPages := 0
	for i := 1; i <= totalPages || totalPages == 0; i++ {
		fmt.Println("page", i)
		resp, err := c.NewRequest().
			WithRegion(official.RegionOceana).
			WithMatchType(official.MatchTypeCustom).
			WithTeamSize(official.TeamSize4v4).
			WithPage(i).
			Do(context.Background()) // use a proper context in an actual app
		if err != nil {
			log.Fatal(err)
		}
		// youd probably add to an array or write to a db here.
		if totalPages == 0 {
			totalPages = (resp.Count / 100) + 1
			fmt.Printf("total records found %d, pages required %d\n", resp.Count, totalPages)
		}
	}
}
