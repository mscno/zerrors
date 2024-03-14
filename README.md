# Zerrors

A error primitives package with a list of grpc based errors that can be used to describe root causes.

Originally forked and extracted from Zitadel `zerrors` package: [Zitadel Errors](https://github.com/zitadel/zitadel/tree/main/internal/zerrors)


## Install

`go get github.com/mscno/zerrors`

This will install the base `zerrors` package and as well as the httperrors adapter.

If you want to use the grpc adapter, you will need to install the `zerrors-grpc` package:

`go get github.com/mscno/zerrors/grpczerrors`

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/mscno/zerrors"
)

func main() {
	rooterr := fmt.Errorf("root error")
    err := zerrors.ThrowNotFound(rooterr,"NOT_FOUND","my error")
    fmt.Println(err)
}
```