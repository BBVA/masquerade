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
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**

Check functions
===============
Since csv is just unttyped format this functions will provide some basic type checking.

*/

// isFloatNumber
func isFloatNumber(x string) bool {
	clean := strings.TrimSpace(x)
	valid := regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
	return valid.MatchString(clean)
}

// isIntNumber
func isIntNumber(x string) bool {
	clean := strings.TrimSpace(x)
	valid := regexp.MustCompile(`^[0-9]+$`)
	return valid.MatchString(clean)
}

// StringToRow convert from raw csv string to internal row
func StringToRow(fieldSeparator string) func(string) ([]interface{}, error) {
	return func(line string) ([]interface{}, error) {
		parts := strings.Split(line, fieldSeparator)
		row := make([]interface{}, len(parts))

		for i, x := range parts {
			clean := strings.TrimSpace(x)
			var field interface{}
			if isFloatNumber(clean) {
				rawFloat, err := strconv.ParseFloat(clean, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cant parse float: %v\n", err)
				}
				field = rawFloat
			} else if isIntNumber(clean) {
				rawInteger, err := strconv.Atoi(clean)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cant parse int: %v\n", err)
				}
				field = rawInteger
			} else {
				rawString := strings.Replace(clean, "\"", "", -1)
				field = rawString
			}
			if field == nil {
				return make([]interface{}, 0), fmt.Errorf("%v cant be formatted to a field", x)
			}
			row[i] = field
		}
		return row, nil
	}
}

// RowToBytes just convert row to bytes
func RowToBytes(fieldSeparator string, lineDelimiter byte, parts int) func(row []interface{}) ([]byte, error) {
	return func(event []interface{}) ([]byte, error) {
		out := []byte{}
		first := true
		for i := 0; i < parts; i++ {
			if !first {
				out = append(out, fieldSeparator...)
			}
			first = false
			var value interface{}
			if i < len(event) {
				value = event[i]
			} else {
				value = nil
			}
			switch v := value.(type) {
			case string:
				out = strconv.AppendQuote(out, v)
			case []uint8:
				out = strconv.AppendQuote(out, string([]byte(v)))
			case int:
				out = strconv.AppendInt(out, int64(v), 10)
			case int64:
				out = strconv.AppendInt(out, v, 10)
			case float64:
				out = strconv.AppendFloat(out, v, 'f', -1, 64)
			default:
				fmt.Fprintf(os.Stderr, "unexpected type %v\n", reflect.TypeOf(value))
			}
		}
		out = append(out, lineDelimiter)
		return out, nil
	}
}
