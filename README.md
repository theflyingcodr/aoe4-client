# AOE4 Client

[![Release](https://img.shields.io/github/release-pre/theflyingcodr/aoe4-client.svg?logo=github&style=flat&v=1)](https://github.com/theflyingcodr/aoe4-client/releases)
[![Build Status](https://github.com/theflyingcodr/aoe4-client/actions/workflows/go.yml/badge.svg)](https://github.com/theflyingcodr/aoe4-client/actions/workflows/go.yml)
[![CodeQL Status](https://github.com/theflyingcodr/aoe4-client/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/theflyingcodr/aoe4-client/actions/workflows/codeql-analysis.yml)
[![Report](https://goreportcard.com/badge/github.com/theflyingcodr/aoe4-client?style=flat&v=1)](https://goreportcard.com/report/github.com/theflyingcodr/aoe4-client)
[![codecov](https://codecov.io/gh/libsv/go-bt/branch/master/graph/badge.svg?v=1)](https://codecov.io/gh/theflyingcodr/aoe4-client)
[![Go](https://img.shields.io/github/go-mod/go-version/theflyingcodr/aoe4-client?v=1)](https://golang.org/)

## Overview

This is a go client used to query AOE4 data from either the official leaderboards or the aoeiv.net api.

There are two packages within the client that can be used [official](official) and [aoeivnet](aoeivnet), each one calls the respective web service but they both contain a similar, fluent filter based API for ease of use.

Under each package is a README explaining useage and with an examples sub directory with code useage.


## Installation

**This** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).

```shell script
go get -u github.com/theflyingcodr/aoe4-client
```

## Documentation

View the generated [documentation](https://pkg.go.dev/github.com/theflyingcodr/aoe4-client)

[![GoDoc](https://godoc.org/github.com/theflyingcodr/aoe4-client?status.svg&style=flat)](https://pkg.go.dev/github.com/theflyingcodr/aoe4-client)

