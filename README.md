# gofigi

Go binding for [OpenFIGI API v3](https://www.openfigi.com/api#v3).

## Overview
GoFiGI provides a Go interface to the [OpenFIGI API](https://www.openfigi.com/api). To access the [OpenFIGI API](https://www.openfigi.com/api)
and benefits from better rate limiting an account and token are required. The goal is for openfigi-go to be compatible with the v3 version of the [OpenFigi API](https://www.openfigi.com/api).

## Motivation
[OpenFIGI](https://www.openfigi.com/api) is a great API with no go-biding.

## Installation
```bash
go get github.com/tochemey-lab/gofigi
```

## Features
- [Mapping API](./openfigi/mapping_client.go)
- [Search API](./openfigi/search_client.go)
- [Filter API](./openfigi/filter_client.go)

## Contribution
Contributions are welcome!
The project adheres to [Semantic Versioning](https://semver.org) and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
This repo uses [Earthly](https://earthly.dev/get-earthly).

To contribute please:
- Fork the repository
- Create a feature branch
- Code 
- Submit a [pull request](https://help.github.com/articles/using-pull-requests)

### Testing and Linter
Prior to submitting a [pull request](https://help.github.com/articles/using-pull-requests), please run:
```bash
earthly +all
```
