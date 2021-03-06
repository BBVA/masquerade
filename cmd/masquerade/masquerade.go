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
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/BBVA/masquerade/pkg/mask"
	"github.com/BBVA/masquerade/pkg/row"

	"github.com/ugorji/go/codec"
)

var (
	fieldsPtr string
)

var rootCommand = &cobra.Command{
	Use:   "masquerade",
	Short: "Masquerade allow you to anonimize data to safely make analitics",
	Run:   masqueradeMain,
}

func masqueradeMain(cmd *cobra.Command, args []string) {
	fields := strings.Split(fieldsPtr, ",")

	var handle codec.Handle = new(codec.MsgpackHandle)
	reader := bufio.NewReader(os.Stdin)
	dec := codec.NewDecoder(reader, handle)

	masker := mask.Factory(fields)
	binFormat := row.Row2Bytes()

	for {
		resMsg := make([]interface{}, len(fields))
		err := dec.Decode(&resMsg)

		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			} else {
				fmt.Fprintf(os.Stderr, "Unexpected error: %v\n", err)
				os.Exit(1)
			}
		}

		row, err := masker(resMsg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to mask: %v\n", resMsg)
		}

		b, err := binFormat(row)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		_, err = os.Stdout.Write(b)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	rootCommand.Flags().StringVar(&fieldsPtr,
		"fields", "",
		"Mask fields separated by , use ,, if no mask and sha256 to mask",
	)
	rootCommand.MarkFlagRequired("fields")
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}
