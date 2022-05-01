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

package iterkit

import (
	"testing"

	ispkg "github.com/matryer/is"
)

func TestSliceIterator(t *testing.T) {
	is := ispkg.New(t)

	it := &SliceIterator[int]{Data: []int{1, 2, 3}}

	var values []int
	for it.Next() {
		values = append(values, it.Value())
	}

	is.Equal(values, []int{1, 2, 3})
}

func BenchmarkSliceIterator(b *testing.B) {
	it := &SliceIterator[int]{}

	for n := 0; n < b.N; n++ {
		it.Next()
	}
}