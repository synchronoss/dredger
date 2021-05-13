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
package converter_logic

import (
	"testing"
)

func TestConverterLogic(t *testing.T) {

	tests := []struct{
		name string
		logic ConverterLogic
		float, string, int, array_string string
	} {
		{
			name: "map",
			logic: CreateMapConverterLogic(),
			float: " 22.000000",
			string: ` "hello"`,
			int: " 22",

			array_string: "key = [\n  hello,\n  there,\n]",
		},
	}

	testOut := func(name string, have string, want string) {
		if have != want {
			t.Errorf("Bad %s format have: %s\nwant: %s", name, have, want)
		}
	}


	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			testOut("float", test.logic.FormatFloat("", 22.0), test.float)
			testOut("string", test.logic.FormatString("", "hello"), test.string)

			testOut("int", test.logic.FormatInt("", 22), test.int)
			testOut("[]string", test.logic.FormatArray("key", []string{"hello", "there"}), test.array_string)
		})

	}
}
