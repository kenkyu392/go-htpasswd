# go-htpasswd

[![test](https://github.com/kenkyu392/go-htpasswd/workflows/test/badge.svg?branch=master)](https://github.com/kenkyu392/go-htpasswd)
[![codecov](https://codecov.io/gh/kenkyu392/go-htpasswd/branch/master/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-htpasswd)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/kenkyu392/go-htpasswd)
[![go report card](https://goreportcard.com/badge/github.com/kenkyu392/go-htpasswd)](https://goreportcard.com/report/github.com/kenkyu392/go-htpasswd)
[![license](https://img.shields.io/github/license/kenkyu392/go-htpasswd)](LICENSE)

Htpasswd parser for Go.

## Installation

```
go get -u github.com/kenkyu392/go-htpasswd
```

## Usage

```go
package main

import (
	"os"

	"github.com/kenkyu392/go-htpasswd"
)

func main() {
	f, err := os.Open(".htpasswd")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := htpasswd.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		// output: name, pass
		println(record[0], record[1])
	}
}
```

## License

[MIT](LICENSE)
