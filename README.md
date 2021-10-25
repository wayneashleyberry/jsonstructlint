> An opinionated linter for json struct tags in Go

[![Build Status](https://travis-ci.org/wayneashleyberry/jsonstructlint.svg?branch=master)](https://travis-ci.org/wayneashleyberry/jsonstructlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonstructlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonstructlint)

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

### Rules

- `json` struct tags must be lower camel case eg. `camelCase`
- `json` struct tags must not contain whitespace
- `json` struct tags must exist on all fields, if they exist on one

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
