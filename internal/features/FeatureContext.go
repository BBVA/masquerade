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

import "github.com/DATA-DOG/godog"

// Feat it's the expected struct to export your steps
type Feat struct {
	expr interface{}
	step interface{}
}

// FillContext fills the step definitions
func FillContext(s *godog.Suite) {
	ctx := &Context{}
	s.BeforeScenario(func(interface{}) {
		ctx.m = make(map[string]interface{})
	})

	for _, feat := range Sha256Features(ctx) {
		s.Step(feat.expr, feat.step)
	}

	for _, feat := range CsvFeatures(ctx) {
		s.Step(feat.expr, feat.step)
	}

	for _, feat := range RabbitFeatures(ctx) {
		s.Step(feat.expr, feat.step)
	}
}
