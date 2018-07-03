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
package features

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/ugorji/go/codec"
)

// Table2Bin converts gherkin table into bin data
func Table2Bin(table *gherkin.DataTable) ([]byte, error) {
	data := gherkinTableToRow(table)
	if len(data) == 0 {
		return nil, fmt.Errorf("Problem decoding test: %+v", data)
	}
	buffer := new(bytes.Buffer)
	var handle codec.Handle = new(codec.MsgpackHandle)
	enc := codec.NewEncoder(buffer, handle)
	for _, dat := range data {
		err := enc.Encode(dat)
		if err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

// BinMsgPackVsTable compares binary to gherkin table
func BinMsgPackVsTable(res []byte, table *gherkin.DataTable) error {
	var handle codec.Handle = new(codec.MsgpackHandle)
	dec := codec.NewDecoderBytes(res, handle)

	expected := gherkinTableToRow(table)

	for _, cur := range expected {
		var resMsg []interface{}
		dec.Decode(&resMsg)
		for i := range resMsg {
			v := resMsg[i]
			intsValue, ok := v.([]uint8)
			if !ok {
				return fmt.Errorf("Expect a string but %v obtained in %v", reflect.TypeOf(res), res)
			}
			strValue := ints2string(intsValue)
			strValue = strings.TrimSpace(strValue)
			r := cur[i]
			if strValue != r {
				s := "[%d] %s != %s\nres:%v\nexpected:%v\ncompleteRes:%v\ncompleteExpected:%v"
				return fmt.Errorf(s, i, r, v, resMsg[i], cur, string(res), expected)
			}

		}
	}
	return nil
}

func ints2string(ints []uint8) string {
	b := make([]byte, len(ints))
	for i, v := range ints {
		b[i] = byte(v)
	}
	return string(b)
}

func gherkinTableToRow(table *gherkin.DataTable) [][]string {
	var out = make([][]string, len(table.Rows))

	for i, x := range table.Rows {
		row := make([]string, len(x.Cells))

		for j := 0; j < len(x.Cells); j++ {
			value := strings.TrimSpace(x.Cells[j].Value)
			row[j] = value
		}
		out[i] = row
	}

	return out
}

func gherkinTableToSlice(fields *gherkin.DataTable) []string {
	out := make([]string, len(fields.Rows))
	var buf bytes.Buffer

	for i, row := range fields.Rows {
		buf.Reset()
		out[i] = row.Cells[0].Value
	}
	return out
}
