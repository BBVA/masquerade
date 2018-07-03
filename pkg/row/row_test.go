// Copyright 2018 BBVA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package row

import "testing"

func TestRowEquals(t *testing.T) {
	rowBuilder := NewRow(-1)
	r1 := rowBuilder("hi", "Tom")
	r2 := rowBuilder("hi", "Tom")

	isEquals := Equals(r1, r2)
	if !isEquals {
		t.Errorf("%v and %v should be equals", r1, r2)
	}
}

func TestAbsentValues(t *testing.T) {
	rowBuilder := NewRow(2)
	res := rowBuilder("hi")

	if res[0] != "hi" {
		t.Errorf("Expected %v but %v obtain", "hi", res[0])
	}
	if res[1] != nil {
		t.Errorf("Expected nil but %v obtain", res[1])
	}
}
