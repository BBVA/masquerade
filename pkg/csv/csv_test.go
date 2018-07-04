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
package csv

import (
	"bytes"
	r "github.com/BBVA/masquerade/pkg/row"
	"testing"
)

func isXTestRun(t *testing.T, inputs []string, expected []bool, f func(string) bool) {
	for i := range inputs {
		inp := inputs[i]
		exp := expected[i]
		res := f(inp)
		if res != exp {
			t.Errorf("when try to detect %v as float get %v but %v expected", inp, res, exp)
		}
	}
}

func TestIsFloat(t *testing.T) {
	inputs := []string{
		"juan",
		"0.0",
		"0.0.0",
		"111111111.11111111",
		"666",
	}
	expected := []bool{
		false,
		true,
		false,
		true,
		false,
	}

	isXTestRun(t, inputs, expected, isFloatNumber)
}

func TestIsInt(t *testing.T) {
	inputs := []string{
		"juan",
		"0",
		"0.0",
		"111111110",
		"666@",
	}
	expected := []bool{
		false,
		true,
		false,
		true,
		false,
	}

	isXTestRun(t, inputs, expected, isIntNumber)
}

func checkError(t *testing.T) func(error, string) {
	return func(err error, msg string) {
		if err != nil {
			t.Errorf(msg+": %v", err)
		}
	}
}

func TestStringToRow(t *testing.T) {
	inputs := []string{
		"\"hi\";\"Bob Dylan\"",
		"\"hi\";666",
		"\"hi\";6.66",
	}

	expected := make([][]interface{}, 3)
	builder := r.NewRow(2)
	expected[0] = builder("hi", "Bob Dylan")
	expected[1] = builder("hi", 666)
	expected[2] = builder("hi", 6.66)

	formatter := StringToRow(";")

	for i := range expected {
		inp := inputs[i]
		exp := expected[i]
		res, err := formatter(inp)

		if !r.Equals(res, exp) {
			t.Errorf("%v expected but %v found", exp, res)
		}
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
	}
}

func TestRowToBytes(t *testing.T) {
	inp := make([][]interface{}, 4)
	builder := r.NewRow(2)
	inp[0] = builder("hi", "Bob Dylan")
	inp[1] = builder("hi", 666)
	inp[2] = builder("hi", 6.66)
	inp[3] = builder([]uint8("hi"), []uint8("Bob Dylan"))
	expected := []string{
		"\"hi\";\"Bob Dylan\"\n",
		"\"hi\";666\n",
		"\"hi\";6.66\n",
		"\"hi\";\"Bob Dylan\"\n",
	}

	formatter := RowToBytes(";", '\n', 2)

	for i := range inp {
		r := inp[i]
		exp := []byte(expected[i])
		res, _ := formatter(r)
		if !bytes.Equal(exp, res) {
			t.Errorf("%v expected but %v found", string(exp), string(res))
		}
	}
}
