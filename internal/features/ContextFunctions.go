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
	"fmt"
	"reflect"
)

// Context has no function, it's intented that only stores Steps
type Context struct {
	m map[string]interface{}
}

func get(m map[string]interface{}, key string) (interface{}, error) {
	storedValue, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found on context", key)
	}
	return storedValue, nil
}

func getStringSlice(m map[string]interface{}, key string) ([]string, error) {
	storedValue, err := get(m, key)
	if err != nil {
		return []string{}, err
	}

	switch v := storedValue.(type) {
	case []string:
		return v, nil
	default:
		return []string{}, fmt.Errorf("unexpected type stored, expected []string but %v found", reflect.TypeOf(v))
	}
}

func getString(m map[string]interface{}, key string) (string, error) {
	storedValue, err := get(m, key)
	if err != nil {
		return "", err
	}

	switch v := storedValue.(type) {
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("unexpected type stored, expected string but %v found", reflect.TypeOf(v))
	}
}

func getByteSlice(m map[string]interface{}, key string) ([]byte, error) {
	storedValue, err := get(m, key)
	if err != nil {
		return []byte{}, err
	}

	switch v := storedValue.(type) {
	case []byte:
		return v, nil
	default:
		return []byte{}, fmt.Errorf("unexpected type stored, expected []byte but %v found", reflect.TypeOf(v))
	}
}

func getShellResult(m map[string]interface{}, key string) (ShellResult, error) {
	storedValue, err := get(m, key)
	if err != nil {
		return ShellResult{}, err
	}

	switch v := storedValue.(type) {
	case ShellResult:
		return v, nil
	default:
		return ShellResult{}, fmt.Errorf("unexpected type stored, expected ShellResult but %v found", reflect.TypeOf(v))
	}
}

func addToParams(m map[string]interface{}, params []string) error {
	_, hasParams := m["params"]

	if hasParams {
		lastParams, err := getStringSlice(m, "params")
		if err != nil {
			return err
		}
		m["params"] = append(lastParams, params...)
	} else {
		m["params"] = params
	}

	return nil
}
