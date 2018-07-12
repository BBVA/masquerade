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

	"github.com/BBVA/masquerade/pkg/rabbit"

	"github.com/DATA-DOG/godog/gherkin"
)

// RabbitFeatures it's Rabbit Features
func RabbitFeatures(ctx *Context) []Feat {
	return []Feat{
		Feat{`^Dial parameter "([^"]*)"$`, ctx.dialParameter},
		Feat{`^No Channel$`, ctx.noChannel},
		Feat{`^Channel "([^"]*)"$`, ctx.channel},
		Feat{`^Channel "([^"]*)" with lines:$`, ctx.channelWithLines},
		Feat{`^Channel "([^"]*)" contains:$`, ctx.channelContains},
	}
}

func (ctx *Context) noChannel() error {
	return nil
}

func (ctx *Context) channel(channel string) error {
	params := make([]string, 2)
	params[0] = "--channel"
	params[1] = channel

	return addToParams(ctx.m, params)
}

func (ctx *Context) dialParameter(dial string) error {
	params := make([]string, 2)
	params[0] = "--dial"
	params[1] = dial

	ctx.m["dial"] = dial
	return addToParams(ctx.m, params)
}

func (ctx *Context) channelWithLines(channel string, lines *gherkin.DocString) error {
	var (
		dial string
		err  error
	)

	ctx.channel(channel)
	linesSlice := strings.Split(lines.Content, "\n")
	linesSize := len(linesSlice)
	params := make([]string, 2)
	params[0] = "--quantity"
	params[1] = fmt.Sprintf("%d", linesSize)
	err = addToParams(ctx.m, params)
	if err != nil {
		return err
	}

	_, hasDial := ctx.m["dial"]

	if hasDial {
		dial, err = getString(ctx.m, "dial")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Expect dial parameter to be in context. %v", ctx.m)
	}

	return rabbit.WriteOnChannel(dial, channel, linesSlice)
}

func (ctx *Context) channelContains(channel string, lines *gherkin.DocString) error {
	linesSlice := strings.Split(lines.Content, "\n")
	nLines := len(linesSlice)

	_, hasDial := ctx.m["dial"]

	var (
		dial string
		err  error
	)

	if hasDial {
		dial, err = getString(ctx.m, "dial")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Expect dial parameter to be in context")
	}

	fromChan, err := rabbit.ReadFromChannel(dial, channel, nLines)
	if err != nil {
		return err
	}

	stdout := strings.Join(fromChan, "\n")

	if !strings.Contains(stdout, lines.Content) {
		return fmt.Errorf("Expect to found %s on output but not found: %+v", lines.Content, stdout)
	}

	return nil
}
