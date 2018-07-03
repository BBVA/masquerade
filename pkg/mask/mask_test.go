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
package mask

import (
	"masquerade/pkg/row"
	"testing"
)

func TestMaskFactory(t *testing.T) {
	conf := []string{
		"sha256",
		"sha256",
		"sha256",
		"",
	}
	f := Factory(conf)
	if f == nil {
		t.Error("Factory must return a function but nil found")
		t.FailNow()
	}

	fields := []string{"f1", "f2", "f3", "f4"}
	r := row.NewRow(4)("Secret", 666, 66.6, "Clear")

	outRow, err := f(r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}
	expected := []string{
		"7e32a729b1226ed1270f282a8c63054d09b26bc9ec53ea69771ce38158dfade8",
		"c7e616822f366fb1b5e0756af498cc11d2c0862edcb32ca65882f622ff39de1b",
		"250a0d04c8b337e4cd4fb29decf082204b543dc3a6f1b9338b5cbe32693c788e",
		"Clear",
	}
	for index := range fields {
		value := outRow[index]
		res, ok := value.(string)
		if !ok || res != expected[index] {
			t.Errorf("expected %v but %v found", expected[index], res)
		}
	}
}
