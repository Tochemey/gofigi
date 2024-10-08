# Contributions are welcome

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
