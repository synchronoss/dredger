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

//
// indent
//
// Indent a block of text by two spaces
//

func Indent(str string) string {
	lines := s.Split(str, "\n")
	var ret []string

	for _, line := range lines {
		if len(line) > 0 {
			ret = append(ret, "  "+line)
		}
	}

	return s.Join(ret, "\n")
}
