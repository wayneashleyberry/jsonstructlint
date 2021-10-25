> An opinionated linter for json struct tags in Go

[![Go](https://github.com/wayneashleyberry/jsonstructlint/actions/workflows/go.yml/badge.svg)](https://github.com/wayneashleyberry/jsonstructlint/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonstructlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonstructlint)

This linter is based on [a post by Fatih Arslan](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/), which is a fantastic read and highly recommended.

### Rules

- `json` struct tags must be lower camel case eg. `camelCase`
- `json` struct tags must not contain whitespace
- `json` struct tags must exist on all fields, if they exist on one

### Example

```sh
go get -u github.com/wayneashleyberry/jsonstructlint
jsonstructlint ./...
test_data/testdata.go:14:2: `x_y` is not camelcase
test_data/testdata.go:16:2: `foo bar` contains whitespace
test_data/testdata.go:17:2: `TitleCase` is not camelcase
test_data/testdata.go:18:2: `a b` contains whitespace
test_data/testdata.go:25:3: `Inline Struct` contains whitespace
test_data/testdata.go:25:3: `Inline Struct` is not camelcase
test_data/testdata.go:30:4: `Super Inline` contains whitespace
test_data/testdata.go:30:4: `Super Inline` is not camelcase
test_data/testdata.go:40:4: `FileName` is not camelcase
test_data/testdata.go:12:2: `F2` is missing a json tag
test_data/testdata.go:13:2: `F3` is missing a json tag
test_data/testdata.go:41:4: `MissingStructTag` is missing a json tag
```

### Editor Config

This linter checks that code conforms to the following vscode config:

```json
"go.addTags": {
  "tags": "json",
  "options": "json=omitempty",
  "promptForTags": false,
  "transform": "camelcase"
}
```
