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
	s "strings"
)

//
// PolicyPath defines the data and logic for a policy path expression.
//
// An exmaple of a policy path is foo.**.baz.*.bar

type PolicyPath interface {
	MatchKey(string) bool
	FullMatch() bool
	PopKey(string) PolicyPath
}

type policyPath struct {
	path []string
}

func (pp policyPath) MatchKey(key string) bool {
	return len(pp.path) > 0 && (pp.path[0] == "*" || pp.path[0] == "**" || pp.path[0] == key)
}

func (pp policyPath) FullMatch() bool {
	return len(pp.path) == 0
}

func (pp policyPath) PopKey(key string) PolicyPath {
	p := pp.path
	if p[0] == "**" {
		if p[1] == key {
			pp.path = p[2:]
		}
	} else {
		pp.path = p[1:]
	}

	return pp
}

func Create(path string) PolicyPath {

	var ret = s.Split(path, ".")
	// If the path is an empty string just return
	// an empty array. Otherise it would match root.
	if len(ret) == 1 && ret[0] == "" {
		ret = []string{}
	}

	return policyPath{path: ret}
}
