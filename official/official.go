package official

import (
	"context"
	"net/http"

	aoe "github.com/theflyingcodr/aoe4-client"
)

// Client is used to query the aoe leaderboards and contains
// an http client.
type Client struct {
	h aoe.HTTPClient
}

// NewClient will setup a new client used for querying the official AOE
// Leaderboards.
func NewClient(c aoe.HTTPClient) *Client {
	return &Client{
		h: c,
	}
}

// ClientRequest allows a fluent filter to be built using
// the Client setup previously. THis allows re-use of the connection pool.
type ClientRequest struct {
	args *LeaderboardArgs
	c    *Client
}

// NewRequest will setup a new client request that can then be filter using
// the fluent With functions.
// To then execute the request call the Do(ctx) method.
// The default settings match the website when first loaded, so if that's what you want
// you can simply call NewRequest().Do(ctx).
func (c *Client) NewRequest() *ClientRequest {
	return &ClientRequest{
		args: &LeaderboardArgs{
			Region:    RegionGlobal,
			Versus:    VersusPlayers,
			MatchType: MatchTypeUnranked,
			TeamSize:  TeamSize1v1,
			Page:      1,
			Count:     100,
		},
		c: c,
	}
}

// Do will execute the request based on the filters supplied before it.
// An error will be returned if we get an unexpected result from the API,
// otherwise the LeaderBoard list will be returned.
func (c *ClientRequest) Do(ctx context.Context) (*Leaderboard, error) {
	var resp *Leaderboard
	return resp, c.c.h.Do(ctx, http.MethodPost, "https://api.ageofempires.com/api/ageiv/Leaderboard", c.args, &resp)
}

// WithVersus will allow you to filter results based on the
// type of matching playing player / ai.
func (c *ClientRequest) WithVersus(v Versus) *ClientRequest {
	c.args.Versus = v
	return c
}

// WithPlayer will filter results down to one player, it seems the API
// is case sensitive so be careful when passing values to this.
func (c *ClientRequest) WithPlayer(p string) *ClientRequest {
	c.args.SearchPlayer = p
	return c
}

// WithMatchType will filter by the type of match, I expect this
// to have more in future. It defaults to unranked if not supplied.
func (c *ClientRequest) WithMatchType(m MatchType) *ClientRequest {
	c.args.MatchType = m
	return c
}

// WithTeamSize will return results filtered for the selected teamsize,
// If not supplied the request isn't filtered by team size.
func (c *ClientRequest) WithTeamSize(t TeamSize) *ClientRequest {
	c.args.TeamSize = t
	return c
}

// WithRegion will return results filtered by region, the default if not
// provided is Global.
func (c *ClientRequest) WithRegion(r Region) *ClientRequest {
	c.args.Region = r
	return c
}

// WithPage will set the current page of results to return, the leaderboard
// api is paged with a max number of 100 per page so you will need to iterate
// until you reach the count return in the LeaderBoard response.
// This defaults to 1 if not supplied.
func (c *ClientRequest) WithPage(p int) *ClientRequest {
	c.args.Page = p
	return c
}

// PageSize is the total amount of records to return for the current page
// the api supports a max of 100, any value greater than this is ignored.
// THis value defaults to 100.
func (c *ClientRequest) PageSize(s int) *ClientRequest {
	c.args.Count = s
	return c
}
