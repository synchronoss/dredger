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
package output_target

import (
	"regexp"
	s "strings"
)

//
// This function tries to create a unique string from a leaf/unit.
//
// e.g.
// resource "kubernetes_statefulset" "foo" => resource_kubernetes_statefulset_foo
//
// This is used to determine the filename when writing output to a directory.
//

var declerationRegex = regexp.MustCompile(`(?m)^(resource|variable|output|module|provider) (:?("\S+"\s*))+{\s*$`)
var alpha = regexp.MustCompile(`(\w+)`)

func DetermineUnitPath(unit string) string {

	match := declerationRegex.FindStringSubmatch(unit)

	if len(match) == 0 {
		return ""
	}

	alphas := alpha.FindAllString(match[0], -1)

	return s.Join(alphas, "_")
}
