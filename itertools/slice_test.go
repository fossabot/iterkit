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
	"testing"

	assertpkg "github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	assert := assertpkg.New(t)

	s := Slice(Range(10))
	assert.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
}
