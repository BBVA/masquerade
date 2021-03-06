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

	"github.com/BBVA/masquerade/pkg/csv"
	"github.com/BBVA/masquerade/pkg/row"
	"github.com/spf13/cobra"
)

var sepPtr string

var rootCmd = &cobra.Command{
	Use:   "maskcsvin",
	Short: "masquerade csv import command",
	Run:   csvInMain,
}

func csvInMain(cmd *cobra.Command, args []string) {
	csvParse := csv.StringToRow(sepPtr)
	binFormat := row.Row2Bytes()

	snr := bufio.NewScanner(os.Stdin)
	for snr.Scan() {
		line := snr.Text()
		if len(line) == 0 {
			break
		}

		row, err := csvParse(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
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

	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	rootCmd.Flags().StringVar(&sepPtr,
		"separator", ",",
		"Separator to use in csv format",
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
