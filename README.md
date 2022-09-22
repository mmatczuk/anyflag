# Anyflag

[![Build Status](https://github.com/mmatczuk/anyflag/actions/workflows/go.yml/badge.svg)](https://github.com/mmatczuk/anyflag/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mmatczuk/anyflag)](https://goreportcard.com/report/github.com/mmatczuk/anyflag)

Anyflag is implementation of [Cobra](https://github.com/spf13/cobra) `Value` and `SliceValue` interfaces using Go Generics.
It allows you to use any type as a flag value.

## Installation

```bash
go get github.com/mmatczuk/anyflag
```

## Usage

```go
    var ba *url.Userinfo
    cmd.Flags().VarP(anyflag.NewValue[*url.Userinfo](nil, &ba, parseBasicAuth), "basic-auth", "", "basic auth")

    var bas []*url.Userinfo
    cmd.Flags().VarP(anyflag.NewSliceValue[*url.Userinfo](nil, &bas, parseBasicAuth), "basic-auth", "", "basic auth")
```

See full running example in [example_test.go](example_test.go).

## License

This project is based on [spf13/pflag](https://github.com/spf13/pflag) licensed under the BSD 3-Clause "New" or "Revised" License
