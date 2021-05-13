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
package policy_set

import (
	"github.com/synchronoss/dredger/components/parsed_leaf"
	"github.com/synchronoss/dredger/components/policy"
)

type (
	Policy = policy.Policy
)

type basicPolicySet struct {
	policies []Policy
}

func (ps basicPolicySet) Execute(docLeaf *DocumentLeaf) (ParsedLeaf, error) {
	for _, p := range ps.policies {
		if p.FullMatch() {
			return p.Execute(docLeaf)
		}
	}
	return parsed_leaf.Create(""), nil
}

func (ps basicPolicySet) Match() bool {
	for _, p := range ps.policies {
		if p.FullMatch() {
			return true
		}
	}
	return false
}

func (ps basicPolicySet) Reduce(key string) PolicySet {
	var newPolicies []Policy

	for _, pol := range ps.policies {
		if pol.MatchKey(key) {
			newPolicies = append(newPolicies, pol.PopKey(key))
		}
	}

	ps.policies = newPolicies
	return ps
}
