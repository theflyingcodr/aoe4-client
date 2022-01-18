package main

import (
	"context"
	"fmt"
	"log"

	aoe "github.com/theflyingcodr/aoe4-client"
	"github.com/theflyingcodr/aoe4-client/official"
)

func main() {
	c := official.NewClient(aoe.NewHTTPClient())

	resp, err := c.NewRequest().
		WithRegion(official.RegionNorthAmerica).
		WithMatchType(official.MatchTypeCustom).
		WithTeamSize(official.TeamSize4v4).
		Do(context.Background()) // use a proper context in an actual app
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total results found %d", resp.Count)
}
