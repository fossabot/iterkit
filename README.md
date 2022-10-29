# iterkit

[![License: APACHE-2.0](https://img.shields.io/badge/license-APACHE--2.0-blue?style=flat-square)](https://www.apache.org/licenses/)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2F0x5a17ed%2Fiterkit.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2F0x5a17ed%2Fiterkit?ref=badge_shield)

Dead simple generic iterator interface.

The package `itertools` is meant to implement iterator functions similar to what python has to offer.


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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2F0x5a17ed%2Fiterkit.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2F0x5a17ed%2Fiterkit?ref=badge_large)