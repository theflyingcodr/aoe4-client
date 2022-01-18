# Official

This is a go api client for the official Age Of Empires 4 [Leaderboard](https://www.ageofempires.com/stats/ageiv/).

[![Release](https://img.shields.io/github/release-pre/theflyingcodr/aoe4-client.svg?logo=github&style=flat&v=1)](https://github.com/theflyingcodr/aoe4-client/releases)
[![Build Status](https://github.com/theflyingcodr/aoe4-client/actions/workflows/go.yml/badge.svg)](https://github.com/theflyingcodr/aoe4-client/actions/workflows/go.yml)
[![CodeQL Status](https://github.com/theflyingcodr/aoe4-client/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/theflyingcodr/aoe4-client/actions/workflows/codeql-analysis.yml)
[![Report](https://goreportcard.com/badge/github.com/theflyingcodr/aoe4-client?style=flat&v=1)](https://goreportcard.com/report/github.com/theflyingcodr/aoe4-client)
[![codecov](https://codecov.io/gh/libsv/go-bt/branch/master/graph/badge.svg?v=1)](https://codecov.io/gh/theflyingcodr/aoe4-client)
[![Go](https://img.shields.io/github/go-mod/go-version/theflyingcodr/aoe4-client?v=1)](https://golang.org/)

It can be used to filter in the same manner as the website and can be useful if you wish to build
go based apps and tools for displaying and working with the official data.


## Useage

A quick example is shown below:

```golang
// Setup and reuse a client, this takes an http client with a configurable timeout
c := official.NewClient(aoe.NewHTTPClient())

// create a request using the client and apply filters
// always call Do to execute the request.
resp, err := c.NewRequest().
    WithRegion(official.RegionNorthAmerica).
    WithMatchType(official.MatchTypeCustom).
    WithTeamSize(official.TeamSize4v4).
    Do(context.Background()) // use a proper context in an actual app
if err != nil {
    log.Fatal(err)
}

// do something with the results
fmt.Printf("total results found %d", resp.Count)
```

## Examples

There are a few examples under the [examples](examples) folder, these show paging, different filters etc.

