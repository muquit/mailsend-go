# Compiling from source

Requires the [Go programming language toolchain](https://golang.org/dl/) and git.
*mailsend-go* uses [Go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

To install the binary:

```
$ go install github.com/muquit/mailsend-go@latest
```

The binary will be installed in `$GOPATH/bin/`.

To compile yourself:

```
$ git clone https://github.com/muquit/mailsend-go.git
$ cd mailsend-go
$ go build .
```

Or for platform-specific builds:

```
$ make linux
$ make mac
$ make windows
```

Run `make help` for all available targets.

To list the packages used:

```
$ go list -m "all"
github.com/muquit/mailsend-go
github.com/muquit/gomail v1.0.2
github.com/muquit/quotedprintable v0.0.0-20250204043250-71206103869d
```