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

/**

Guard functions
=================

This functions alow to omit some data and don't break any thing.

*/

// defaulLineDelimiterGuard will put \n as delimiter if it's not provided
func defaulLineDelimiterGuard(lineDelimiter string) byte {
	var ld byte
	if lineDelimiter != "" {
		bytes := []byte(lineDelimiter)
		ld = bytes[0]
	} else {
		ld = '\n'
	}
	return ld
}

// defaultFieldDelimiterGuard will put , as delimiter if it's not provided
func defaultFieldDelimiterGuard(fieldDelimiter string) string {
	var fd string
	if fieldDelimiter != "" {
		fd = fieldDelimiter
	} else {
		fd = ","
	}
	return fd
}
