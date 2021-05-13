/*
 * Copyright (C) 2021 Synchronoss Technologies
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package terraform_format

import (
	"testing"
)

type Test struct {
	original     string
	snake        string
	indent       string
	quote        string
	strip        string
	formatString string
	formatKey    string
}

func RunTests(t *testing.T, cbk func(string) string, extract func(Test) string) {

	tests := []Test{
		{original: "fooBar", snake: "foo_bar", indent: "  fooBar", quote: `"fooBar"`},
		{original: "bazBar", snake: "baz_bar", indent: "  bazBar", quote: `"bazBar"`},
		{original: `"hello"`, quote: `"\"hello\""`},
		{original: "${var.foo}bar", strip: "bar"},
		{original: "${var.foo}-baz", strip: "baz", formatKey: "baz", formatString: `"${var.foo}-baz"`},
		{original: "${var.foo}", formatString: "var.foo"},
		{original: "foo", formatString: `"foo"`},
		{original: "${var.foo}\n${var.bar}", formatString: "<<-EOF\n  ${var.foo}\n  ${var.bar}\n  EOF\n"},
		{original: "${var.name}-foo-bar", formatKey: "foo_bar"},
	}

	for _, test := range tests {
		expect := extract(test)

		if expect == "" {
			continue
		}

		t.Run(test.original, func(t *testing.T) {

			result := cbk(test.original)

			if expect != result {
				t.Errorf("Conversion failed\nwant: %s\nhave: %s", expect, result)
			}
		})
	}

}

func TestSnakeCase(t *testing.T) {
	RunTests(t, SnakeCase, func(test Test) string { return test.snake })
}

func TestIndent(t *testing.T) {
	RunTests(t, Indent, func(test Test) string { return test.indent })
}

func TestQuote(t *testing.T) {
	RunTests(t, Quote, func(test Test) string { return test.quote })
}

func TestStripVars(t *testing.T) {
	RunTests(t, StripVars, func(test Test) string { return test.strip })
}

func TestFormatString(t *testing.T) {
	RunTests(t, FormatStringValue, func(test Test) string { return test.formatString })
}

func TestFormatKey(t *testing.T) {
	RunTests(t, FormatKey, func(test Test) string { return test.formatKey })
}
