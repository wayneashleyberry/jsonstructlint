> An opinionated linter for json struct tags in Go

[![Build Status](https://travis-ci.org/wayneashleyberry/jsonlint.svg?branch=master)](https://travis-ci.org/wayneashleyberry/jsonlint)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/jsonlint)](https://goreportcard.com/report/github.com/wayneashleyberry/jsonlint)

```sh
go get -u github.com/wayneashleyberry/jsonlint
jsonlint ./...
./tesdata/testdata.go:8: "x_y" is not camelcase
./tesdata/testdata.go:10: "foo bar" is not camelcase
./tesdata/testdata.go:11: "TitleCase" is not camelcase
./tesdata/testdata.go:12: "a b" is not camelcase
```
