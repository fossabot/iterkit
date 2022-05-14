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

// Slice consumes the iterator iterkit.Iterator returning its
// elements as a Go slice.
func Slice[T any](it iterkit.Iterator[T]) (out []T) {
	for it.Next() {
		out = append(out, it.Value())
	}
	return
}
