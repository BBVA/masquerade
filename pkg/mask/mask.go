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
	"crypto/sha256"
	"fmt"
	"strconv"
)

func hashSha256(value interface{}) interface{} {
	var aStringToHash string

	switch v := value.(type) {
	case string:
		aStringToHash = v
	case []uint8:
		aStringToHash = string([]byte(v))
	case int:
		aStringToHash = strconv.Itoa(v)
	case int64:
		aStringToHash = strconv.FormatInt(v, 10)
	case float64:
		aStringToHash = strconv.FormatFloat(v, 'f', -1, 64)
	}

	sm := fmt.Sprintf("%x", sha256.Sum256([]byte(aStringToHash)))
	return string(sm)
}

func identity(value interface{}) interface{} {
	return value
}

// Factory generates the function that convert the row
// Maybe for security must fail if field is not defined to avoid
// accidental leaks
func Factory(masks []string) func([]interface{}) ([]interface{}, error) {
	masker := make([]func(interface{}) interface{}, len(masks))

	//TODO: Memoize field values
	for i, method := range masks {
		switch method {
		case "sha256":
			masker[i] = hashSha256
		default:
			masker[i] = identity
		}
	}

	return func(current []interface{}) ([]interface{}, error) {

		row := make([]interface{}, len(masks))

		for i, value := range current {
			row[i] = masker[i](value)
		}

		return row, nil
	}
}
