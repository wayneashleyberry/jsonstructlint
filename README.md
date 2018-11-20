> An opinionated linter for json struct tags in Go

[![Build Status](https://travis-ci.org/wayneashleyberry/jsonstructlint.svg?branch=master)](https://travis-ci.org/wayneashleyberry/jsonstructlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonstructlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonstructlint)

```sh
go get -u github.com/wayneashleyberry/jsonstructlint
jsonstructlint ./...
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:25:3: "Inline Struct" contains whitespace
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:30:4: "Super Inline" contains whitespace
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:14:2: "x_y" is not camelcase
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:16:2: "foo bar" contains whitespace
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:17:2: "TitleCase" is not camelcase
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:18:2: "a b" contains whitespace
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:12:2: F2 is missing a struct tag
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:13:2: F3 is missing a struct tag
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:40:4: "FileName" is not camelcase
/Users/wayneberry/go/src/github.com/wayneashleyberry/jsonstructlint/tesdata/testdata.go:41:4: MissingStructTag is missing a struct tag
```

### Rules

- `json` struct tags must be lower camel case eg. `camelCase`
- `json` struct tags must not contain whitespace
- `json` struct tags must exist on all fields, if they exist on one
