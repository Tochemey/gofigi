# gofigi

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Tochemey/gofigi/build.yml)]((https://github.com/Tochemey/gofigi/actions/workflows/build.yml))
[![codecov](https://codecov.io/gh/Tochemey/gofigi/graph/badge.svg?token=vbFBbfgcrD)](https://codecov.io/gh/Tochemey/gofigi)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tochemey/gofigi)](https://go.dev/doc/install)

Go binding for [OpenFIGI API v3](https://www.openfigi.com/api#v3).

## Overview
`gofigi` provides a Go interface to the [OpenFIGI API](https://www.openfigi.com/api). To access the [OpenFIGI API](https://www.openfigi.com/api)
and benefits from better rate limiting an account and token are required. The goal is for `gofigi` to be compatible with the v3 version of the [OpenFigi API](https://www.openfigi.com/api).

## Motivation
[OpenFIGI API](https://www.openfigi.com/api) is a great API with no go-binding.

## Installation
```bash
go get github.com/tochemey/gofigi
```

## Features
- [Mapping API](https://www.openfigi.com/api#post-v3-mapping)
- [Search API](https://www.openfigi.com/api#post-v3-search)
- [Filter API](https://www.openfigi.com/api#post-v3-filter)
- [Rate Limit](https://www.openfigi.com/api#rate-limit) is built-in all the various APIs

## Contribution
Contributions are welcome!
The project adheres to [Semantic Versioning](https://semver.org) and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

To contribute please:
- Fork the repository
- Create a feature branch
- Code 
- Submit a [pull request](https://help.github.com/articles/using-pull-requests)

### Testing and Linter
Prior to submitting a [pull request](https://help.github.com/articles/using-pull-requests), please run:
```bash
go test -mod=vendor ./... -race -coverprofile=coverage.out -covermode=atomic -coverpkg=./...
```
