# Compiling

Compiling from scratch requires the [Go programming language toolchain](https://golang.org/dl/) and git. Note: *mailsend-go* uses [go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

To install the binary:

```
    $ go install github.com/muquit/mailsend-go@latest
```
The binary will be installed at $GOPATH/bin/ directory.


If you see the error message `go: cannot find main module; see 'go help
modules'`, make sure GO111MODULE environment variable is not set to on. Unset it by
typing `unset GO111MODULE`


To compile yourself:

```
    $ git clone https://github.com/muquit/mailsend-go.git
    $ cd mailsend-go
    $ make
    $ ./mailsend-go -V
```

* List the packages used (if you are outside $GOPATH)
```
    $ go list -m "all"
    github.com/muquit/mailsend-go
    github.com/muquit/gomail v0.0.0-20250327010414-6846ede5e07d
    github.com/muquit/quotedprintable v0.0.0-20250204043250-71206103869d
```
