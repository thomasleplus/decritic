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

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## Security

Please read [SECURITY.md](SECURITY.md) for details on our security policy and how to report security vulnerabilities.

## Code of Conduct

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct.

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
