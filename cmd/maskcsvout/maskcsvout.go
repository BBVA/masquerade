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
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"github.com/BBVA/masquerade/pkg/csv"
	"os"
	"strings"

	"github.com/ugorji/go/codec"
)

func parseFields(fields string) []string {
	return strings.Split(fields, ",")
}

func main() {
	fieldsPtr := flag.Int("fields", 0, "number of fields must have the output")
	sepPtr := flag.String("separator", ",", "field separator")
	flag.Parse()

	//TODO: fields size must be known

	var (
		ld        byte = '\n'
		formatter func(row []interface{}) ([]byte, error)
	)

	var handle codec.Handle = new(codec.MsgpackHandle)
	reader := bufio.NewReader(os.Stdin)
	dec := codec.NewDecoder(reader, handle)
	for {
		resMsg := make([]interface{}, *fieldsPtr)
		err := dec.Decode(&resMsg)

		if formatter == nil {
			if *fieldsPtr == 0 {
				formatter = csv.RowToBytes(*sepPtr, ld, len(resMsg))
			} else {
				formatter = csv.RowToBytes(*sepPtr, ld, *fieldsPtr)
			}
		}

		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			} else {
				fmt.Fprintf(os.Stderr, "Unexpected error: %v\n", err)
				os.Exit(1)
			}
		}

		s, err := formatter(resMsg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unexpected error: %v\n", err)
		} else {
			fmt.Fprint(os.Stdout, string(s))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unexpected error: %v\n", err)
			}
		}
	}
}
