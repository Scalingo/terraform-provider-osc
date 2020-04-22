Terraform Provider for Outscale (unofficial)
==================

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12.20
- [Go](https://golang.org/doc/install) 1.13 at least (to build the provider plugin)

Build and Install
---------------------

Compile the binary and install it:
```go
go get -v
go install
```

Build with Docker (not tested)
---------------------

Build the Docker image

```
$ make docker-image
```

Build the binaries, you'll find all the binaries in pkg/

```
$ make docker-build
```
