> An opinionated linter for json struct tags in Go

[![Build Status](https://travis-ci.org/wayneashleyberry/jsonstructlint.svg?branch=master)](https://travis-ci.org/wayneashleyberry/jsonstructlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonstructlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonstructlint)

```sh
go get -u github.com/wayneashleyberry/jsonstructlint
jsonstructlint ./...
testdata/testdata.go:25:3: "Inline Struct" contains whitespace
testdata/testdata.go:30:4: "Super Inline" contains whitespace
testdata/testdata.go:14:2: "x_y" is not camelcase
testdata/testdata.go:16:2: "foo bar" contains whitespace
testdata/testdata.go:17:2: "TitleCase" is not camelcase
testdata/testdata.go:18:2: "a b" contains whitespace
testdata/testdata.go:12:2: F2 is missing a struct tag
testdata/testdata.go:13:2: F3 is missing a struct tag
testdata/testdata.go:40:4: "FileName" is not camelcase
testdata/testdata.go:41:4: MissingStructTag is missing a struct tag
```

### Rules

- `json` struct tags must be lower camel case eg. `camelCase`
- `json` struct tags must not contain whitespace
- `json` struct tags must exist on all fields, if they exist on one
