# is-docker

> Check if the process is running inside a Docker container

## Install

```sh
$ go get github.com/davipatricio/is-docker
```

## Usage

```go
package main

import (
    "os"
    "github.com/davipatricio/is-docker"
)

func main() {
    if isdocker.Check() {
        fmt.Println("Running inside a Docker container")
    } else {
        fmt.Println("Running outside a Docker container")
        // Exit the process
        os.Exit(1)
    }
}
```

## Related

[is-docker](https://github.com/sindresorhus/is-docker) - Check if the process is running inside a Docker container