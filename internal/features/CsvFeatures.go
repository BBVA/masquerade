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
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
)

// CsvFeatures it's Csv Features
func CsvFeatures(ctx *Context) []Feat {
	return []Feat{
		Feat{`^separator "([^"]*)"$`, ctx.separator},
		Feat{`^pass thru StdIn lines:$`, ctx.passThruStdInLines},
		Feat{`^StdOut should contain lines:$`, ctx.stdOutShouldContainLines},
	}
}

func (ctx *Context) separator(sep string) error {
	params := make([]string, 2)
	params[0] = "--separator"
	params[1] = sep

	return addToParams(ctx.m, params)
}

func (ctx *Context) passThruStdInLines(lines *gherkin.DocString) error {
	ctx.m["stdin"] = []byte(strings.TrimSpace(lines.Content))
	return nil
}

func (ctx *Context) stdOutShouldContainLines(lines *gherkin.DocString) error {
	result, err := getShellResult(ctx.m, "result")
	if err != nil {
		return err
	}
	if !strings.Contains(string(result.Stdout), lines.Content) {
		return fmt.Errorf("Expected %v obtained %v", lines.Content, string(result.Stdout))
	}
	return nil
}
