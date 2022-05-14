// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package itertools

import (
	"github.com/0x5a17ed/iterkit"
)

type EachFunc[T any] func(item T) bool

// Each walks through the iterator given in it and calls fn for
// every single entry.
func Each[T any](it iterkit.Iterator[T], fn EachFunc[T]) {
	for it.Next() {
		if fn(it.Value()) {
			break
		}
	}
}

type EachIndexFunc[T any] func(i int, item T) bool

// EachIndex walks through the iterator given in it and calls fn for
// every single entry together with its index.
func EachIndex[T any](it iterkit.Iterator[T], fn EachIndexFunc[T]) {
	for i := 0; it.Next(); i += 1 {
		if fn(i, it.Value()) {
			break
		}
	}
}
