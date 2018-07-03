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

import (
	"bytes"
	"fmt"

	"github.com/ugorji/go/codec"
)

// Equals check if two rows are the same
func Equals(r1, r2 []interface{}) bool {
	sizeR1 := len(r1)
	sizeR2 := len(r2)
	if sizeR1 != sizeR2 {
		return false
	}

	for index, current := range r1 {
		other := r2[index]
		if other != current {
			return false
		}
	}

	return true
}

// FieldsGuard generate more fields if needed
func FieldsGuard(fields []string, required int) []string {
	fieldsSize := len(fields)
	var validFields []string
	if fieldsSize < required {
		current := 1
		i := 0
		validFields = make([]string, required)
		for i < required {
			if i < fieldsSize {
				validFields[i] = fields[i]
			} else {
				validFields[i] = fmt.Sprintf("field%d", current)
				current++
			}
			i++
		}
	} else {
		validFields = fields
	}
	return validFields
}

// NewRow help us creating a valid Row
// if you pass as fixedFields -1 you are setting a variable row size
func NewRow(fixedFields int) func(...interface{}) []interface{} {

	return func(values ...interface{}) []interface{} {
		var size int
		if fixedFields == -1 {
			size = len(values)
		} else {
			size = fixedFields
		}
		m := make([]interface{}, size)
		for i := 0; i < size; i++ {
			if i < len(values) {
				m[i] = values[i]
			} else {
				m[i] = nil
			}
		}
		return m
	}
}

//Row2Bytes will transform row into its mesgpack byte representation
func Row2Bytes() func([]interface{}) ([]byte, error) {
	//TODO: test this
	handle := new(codec.MsgpackHandle)

	return func(row []interface{}) ([]byte, error) {
		buffer := new(bytes.Buffer)
		enc := codec.NewEncoder(buffer, handle)
		err := enc.Encode(row)
		if err != nil {
			return []byte{}, err
		}
		return buffer.Bytes(), nil
	}
}
