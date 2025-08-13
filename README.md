# Decritic

Remove diacritics (accents)

[![Go](https://github.com/thomasleplus/decritic/workflows/Go/badge.svg)](https://github.com/thomasleplus/decritic/actions?query=workflow:"Go")
[![CodeQL](https://github.com/thomasleplus/decritic/workflows/CodeQL/badge.svg)](https://github.com/thomasleplus/decritic/actions?query=workflow:"CodeQL")

## Installation

```shell
go install github.com/thomasleplus/decritic@latest
```

Make sure that your Golang path (meaning the output of the command `go env GOPATH`, e.g. "${HOME}/go/bin") is included in your PATH environment variable.

## Usage

```shell
cat input.txt | decritic > output.txt
```
