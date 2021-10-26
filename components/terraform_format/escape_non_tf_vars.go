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
// EscapeNonTfVars
//
// Escape any non-terraform variables

var matchNonTfVar = regexp.MustCompile("(\\${[^}]+})")

func escapePrefix(v string) string {
	if s.HasPrefix(v, "${var.") || s.HasPrefix(v, "${local.") {
		return v
	} else {
		return "$$" + EscapeNonTfVars(s.TrimPrefix(v, "$"))
	}
}

func EscapeNonTfVars(v string) string {
	return matchNonTfVar.ReplaceAllStringFunc(v, escapePrefix)
}
