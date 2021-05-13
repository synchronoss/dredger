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
package policy_path

import (
	"testing"
)

func TestCreate(t *testing.T) {

	testStrings := []string{"foo.**.bar.*.baz"}

	for _, testString := range testStrings {
		t.Run(testString, func(t *testing.T) {
			path := Create(testString)

			if path == nil {
				t.Fatalf("Create returns a null for string")
			}
		})

	}
}

func TestMatch(t *testing.T) {

	match := []string{"foo", "bar", "baz"}

	tests := []struct {
		path  string
		count int
	}{
		{path: "foo.bar.baz", count: 3},
		{path: "foo.*.baz", count: 3},
		{path: "**.baz", count: 3},
		{path: "foo.bar", count: 2},
		{path: "foo", count: 1},
		{path: "bar", count: 0},
		{path: "foo.**.baz", count: 3},
		{path: "**.*.baz", count: 3},
		{path: "*.**.baz", count: 3},
	}

	for _, test := range tests {

		t.Run(test.path, func(t *testing.T) {

			var path = Create(test.path)

			var count = 0

			for _, m := range match {

				if path.FullMatch() {
					break
				} else if path.MatchKey(m) {
					path = path.PopKey(m)
					count += 1
				} else {
					break
				}
			}

			if count != test.count {
				t.Fatalf("Path %s has incorrect match. Got %d expected %d", test.path, count, test.count)
			}

		})

	}
}
