# Compiling

Compiling from scratch requires the [Go programming language toolchain](https://golang.org/dl/) and git.

To download, build and install (or upgrade) mailsend-go, run:

```
go get -u github.com/muquit/mailsend-go
```

To compile yourself:

* If you are using very old version of go, install dependecies by typing:

```
make tools
make
```

* If you are using go 1.11+, dependencies will be installed via go modules.
If you cloned mailsend-go inside your $GOPATH, you have to set env var:

```
export GO111MODULE=on
```
* Finally compile mailsend-go by typing:

```
make
```

Type `make help` for more targets:

