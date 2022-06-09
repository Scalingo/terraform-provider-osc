Terraform Provider for Outscale (unofficial)
==================


Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12.20
- [Go](https://golang.org/doc/install) 1.13 at least (to build the provider plugin)

Build and Install
---------------------

Compile the binary and install it:
```shell
go get -v
go install
mkdir -p ~/.terraform.d/plugins/linux_amd64
cp $GOPATH/bin/terraform-provider-osc ~/.terraform.d/plugins/linux_amd64
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

## Release a New Version

Instructions on this [Notion page](https://www.notion.so/scalingooriginal/New-Terraform-Provider-Release-40cd0af66b1f48148fb641ea138a22e5).
