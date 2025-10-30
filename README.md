# Decritic

Remove diacritics (accents)

[![Go](https://github.com/thomasleplus/decritic/workflows/Go/badge.svg)](https://github.com/thomasleplus/decritic/actions?query=workflow:"Go")

## Installation

```shell
go install github.com/thomasleplus/decritic@latest
```

Make sure that your Go path (meaning the output of the command `go env GOPATH`, e.g. "${HOME}/go/bin") is included in your PATH environment variable.

## Usage

To get help, run the following:

```shell
decritic -h
```

In short, `decritic` can either read from files or from `stdin`:

```shell
cat input.txt | decritic > output.txt
decritic input.txt > output.txt
```
