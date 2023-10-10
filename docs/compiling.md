# Compiling

Compiling from scratch requires the [Go programming language toolchain](https://golang.org/dl/) and git. Note: *mailsend-go* uses [go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

To download, build and install (or upgrade) mailsend-go, run:

```
    $ go get -u github.com/muquit/mailsend-go
```
If you see the error message `go: cannot find main module; see 'go help
modules'`, make sure GO111MODULE environment variable is not set to on. Unset it by
typing `unset GO111MODULE`


To compile yourself:

* If you are using very old version of go, install dependencies by typing:

```
    $ make tools
    $ make
```

* If you are using go 1.11+, dependencies will be installed via go modules.
If you cloned mailsend-go inside your $GOPATH, you have to set env var:

```
    $ export GO111MODULE=on
```
* Finally compile mailsend-go by typing:

```
    $ make
```

As mailsend-go uses go modules, it can be built outside $GOPATH e.g.
```
    $ cd /tmp
    $ git clone https://github.com/muquit/mailsend-go.git
    $ cd mailsend-go
    $ make
    $ ./mailsend-go -V
    @(#) mailsend-go v1.0.1
```
* List the packages used (if you are outside $GOPATH)
```
    $ go list -m "all"
    github.com/muquit/mailsend-go
    gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc
    gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
```
Type `make help` for more targets:
