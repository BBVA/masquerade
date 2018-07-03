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

// Sha256Features are the Sha256 Features
func Sha256Features(ctx *Context) []Feat {
	return []Feat{
		Feat{`^No parameters$`, ctx.noParameters},
		Feat{`^Fields:$`, ctx.fields},
		Feat{`^pass thru StdIn msgpack:$`, ctx.passThruStdInMsgpack},
		Feat{`^Invoke "([^"]*)"$`, ctx.invoke},
		Feat{`^exit code must be (-?\d+)$`, ctx.exitCode},
		Feat{`^StdOut should be msgpack:$`, ctx.stdOutShouldBeMsgPack},
		Feat{`^StdOut should be msgpack:$`, ctx.stdOutShouldBeMsgPack},
		Feat{`^Error message should contain "([^"]*)"$`, ctx.errorMessageShouldBe},
	}

}

func (Context) noParameters() error {
	return nil
}

func (ctx *Context) fields(fs *gherkin.DataTable) error {
	params := make([]string, 2)
	params[0] = "-fields"
	fields := gherkinTableToSlice(fs)
	params[1] = strings.Join(fields, ",")

	return addToParams(ctx.m, params)
}

func (ctx *Context) invoke(executable string) error {
	var (
		params []string
		stdin  []byte
		err    error
	)
	_, hasParams := ctx.m["params"]
	if hasParams {
		params, err = getStringSlice(ctx.m, "params")
		if err != nil {
			return err
		}
	}
	_, hasStdin := ctx.m["stdin"]
	if hasStdin {
		stdin, err = getByteSlice(ctx.m, "stdin")
		if err != nil {
			return err
		}
	}

	res, err := CallShell(executable, params, stdin)
	if err != nil {
		return err
	}
	ctx.m["result"] = res
	return nil
}

func (ctx *Context) passThruStdInMsgpack(table *gherkin.DataTable) error {
	bin, err := Table2Bin(table)
	if err != nil {
		return err
	}
	ctx.m["stdin"] = bin
	return nil
}

func (ctx *Context) exitCode(exitArg int) error {
	result, err := getShellResult(ctx.m, "result")
	if err != nil {
		return err
	}
	if result.ExitCodeObtained != exitArg {
		return fmt.Errorf("expected exit code %d but %d obtain: %+v", exitArg, result.ExitCodeObtained, result)
	}
	return nil
}

func (ctx *Context) errorMessageShouldBe(message string) error {
	result, err := getShellResult(ctx.m, "result")
	if err != nil {
		return err
	}
	if !strings.Contains(string(result.Stderr), message) {
		return fmt.Errorf("Expect to found %s on error but not found: %+v", message, result)
	}

	return nil
}

func (ctx *Context) stdOutShouldBeMsgPack(table *gherkin.DataTable) error {
	result, err := getShellResult(ctx.m, "result")
	if err != nil {
		return err
	}

	return BinMsgPackVsTable(result.Stdout, table)
}
