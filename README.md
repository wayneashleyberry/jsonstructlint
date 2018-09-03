> An opinionated linter for json struct tags in Go

[![Build Status](https://travis-ci.org/wayneashleyberry/jsonstructlint.svg?branch=master)](https://travis-ci.org/wayneashleyberry/jsonstructlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonstructlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonstructlint)

```sh
go get -u github.com/wayneashleyberry/jsonstructlint
jsonstructlint ./...
tesdata/testdata.go:8:2: "x_y" is not camelcase
tesdata/testdata.go:10:2: "foo bar" contains whitespace
tesdata/testdata.go:11:2: "TitleCase" is not camelcase
tesdata/testdata.go:12:2: "a b" contains whitespace
tesdata/testdata.go:19:3: "Inline Struct" contains whitespace
tesdata/testdata.go:24:4: "Super Inline" contains whitespace
```
