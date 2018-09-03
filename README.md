> An opinionated linter for json struct tags in Go

[![Build Status](https://travis-ci.org/wayneashleyberry/jsonstructlint.svg?branch=master)](https://travis-ci.org/wayneashleyberry/jsonstructlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonstructlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonstructlint)

```sh
go get -u github.com/wayneashleyberry/jsonstructlint
jsonstructlint ./...
tesdata/testdata.go:14:2: "x_y" is not camelcase
tesdata/testdata.go:16:2: "foo bar" contains whitespace
tesdata/testdata.go:17:2: "TitleCase" is not camelcase
tesdata/testdata.go:18:2: "a b" contains whitespace
tesdata/testdata.go:12:2: Thing.F2 is missing a struct tag
tesdata/testdata.go:13:2: Thing.F3 is missing a struct tag
tesdata/testdata.go:25:3: "Inline Struct" contains whitespace
tesdata/testdata.go:30:4: "Super Inline" contains whitespace
```

### Rules

- `json` struct tags must be lower camel case eg. `camelCase`
- `json` struct tags may not contain whitespace
- If a single struct field has a `json` tag, all fields must
