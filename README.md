
<p align="center">
<img src="https://i.imgur.com/Rlk1t0H.png" alt="Dora" />
</p>


# `Dora`


high-performance Reactive HTTP server based on gnet


# Getting Started

## Features

* High-Performance. Non-Blocking. Asynchronous        
* Simple Code base. High Optimized Routing. Scales on multi-core CPUs
* Extremly low memory usage
* Zero GC overhead
* Simple, pure Go implementation

## Installing

To start using `Dora`, install Go and run `go get`:

```sh
$ go get -u github.com/danyalmh/Dora
```

This will retrieve the library.

## Usage
```go
package main

import (
	"dora.com/web/framework/dorahttp"
)

func main() {

	routing := dorahttp.NewRouter()

	routing.GET("/dora/reactive", func(dctx *dorahttp.Dctx) []byte {

		return dctx.Response(dorahttp.StatusOK, "hello Reactive World ....")
	})


	// Start(Port, MultiCore, Router)
	dorahttp.Start(8089, true, routing)
}


```


## Contact

DanyalMh [@danyalmh](https://github.com/danyalmh)
