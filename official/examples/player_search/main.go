package main

import (
	"context"
	"fmt"
	"log"

	aoe "github.com/theflyingcodr/aoe4-client"
	"github.com/theflyingcodr/aoe4-client/official"
)

// This demonstrates a search for a single player, in this case
// filtering for their global AIHard 2v2 match rankings.
func main() {
	c := official.NewClient(aoe.NewHTTPClient())

	resp, err := c.NewRequest().
		WithRegion(official.RegionGlobal).
		WithMatchType(official.MatchTypeAIHard).
		WithTeamSize(official.TeamSize2v2).
		WithPlayer("DiabolikMark").
		Do(context.Background()) // use a proper context in an actual app
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
