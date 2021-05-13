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
	"regexp"
	s "strings"
)

//
// FormatStringValue
//
// - unwrap interpolation only strings
// - escape non-terraform vars
// - quote of not var-only
//

func isMultiLine(str string) bool {
	lines := s.Split(str, "\n")
	return len(lines) > 1
}

var interpolationOnly = regexp.MustCompile(`^\${(var\.[^}]+)}$`)

func FormatStringValue(v string) string {
	v = EscapeNonTfVars(v)
	if interpolationOnly.MatchString(v) {
		return interpolationOnly.ReplaceAllString(v, "$1")
	} else if isMultiLine(v) {
		return "<<-EOF\n" + Indent(v) + "\n  EOF" + "\n"
	} else {
		return Quote(v)
	}
}
