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
	"github.com/synchronoss/dredger/components/config"
	"github.com/synchronoss/dredger/components/document_leaf"
	"github.com/synchronoss/dredger/components/parsed_leaf"
	"github.com/synchronoss/dredger/components/policy"
)

//
// PolicySet
//

type (
	DocumentLeaf = document_leaf.DocumentLeaf
	ParsedLeaf   = parsed_leaf.ParsedLeaf
	Config       = config.Config
	FuncMap      = policy.FuncMap
)

type PolicySet interface {
	Reduce(string) PolicySet
	Match() bool
	Execute(*DocumentLeaf) (ParsedLeaf, error)
}

func CreatePolicySet(policies []policy.Policy) PolicySet {
	return basicPolicySet{policies: policies}
}

func CreateEmptyPolicySet() PolicySet {
	return basicPolicySet{policies: []policy.Policy{}}
}

func BuildPolicies(config *Config, funcMap *FuncMap) (PolicySet, error) {
	var policies []Policy

	for _, policyConfig := range config.Policies {

		policy, policyErr := policy.BuildPolicy(&policyConfig, *funcMap)

		if policyErr != nil {
			return nil, policyErr
		}

		policies = append(policies, policy)
	}

	return basicPolicySet{policies: policies}, nil
}
