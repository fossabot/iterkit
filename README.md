# iterkit

[![License: APACHE-2.0](https://img.shields.io/badge/license-APACHE--2.0-blue?style=flat-square)](https://www.apache.org/licenses/)

Dead simple generic iterator interface.

## 📦 Installation

```shell
$ go get -u github.com/0x5a17ed/iterkit@latest
```


## 🤔 Usage

```go
package main

import (
	"github.com/0x5a17ed/iterkit"
)

func main() {
	it := &iterkit.SliceIterator[int]{Data: []int{1, 2, 3}}
	for it.Next() {
		fmt.Println(it.Value)()
	}
}
```
