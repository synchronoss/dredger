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
	s "strings"
)

// AssignKey
//
// Set a key to be "key = " as long as key is not empty.

func AssignKey(k string) string {
	if k == "" || k == `""` {
		return ""
	} else {
		return k + " ="
	}
}

func FormatKey(str string) string {
	stripped := StripVars(str)
	underscored := s.Replace(stripped, "-", "_", -1)
	return SnakeCase(underscored)
}
