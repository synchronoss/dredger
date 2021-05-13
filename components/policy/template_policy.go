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
package policy

import (
	"github.com/synchronoss/dredger/components/config"
	"github.com/synchronoss/dredger/components/parsed_leaf"
	"github.com/synchronoss/dredger/components/policy_path"
	"github.com/synchronoss/dredger/components/template"
	s "strings"
)

type Template = template.Template
type FuncMap = template.FuncMap

type templatePolicy struct {
	tpl      Template
	unitTpls []Template
	path     policy_path.PolicyPath
}

func (tp templatePolicy) Execute(leaf *DocumentLeaf) (ParsedLeaf, error) {
	leafString, err := tp.tpl.Execute(leaf)
	if err != nil {
		return parsed_leaf.Create(""), err
	}

	var ret = parsed_leaf.Create(leafString)

	for _, unitTpl := range tp.unitTpls {
		unitString, unitErr := unitTpl.Execute(leaf)
		if unitErr != nil {
			return ret, unitErr
		}

		// units may return empty so trim and check if we should
		// ignore it
		trimmedString := s.Trim(unitString, "\n \t")
		if trimmedString != "" {
			ret = ret.AddUnits(trimmedString + "\n")
		}
	}

	return ret, nil
}

func (tp templatePolicy) MatchKey(key string) bool {
	return tp.path.MatchKey(key)
}

func (tp templatePolicy) FullMatch() bool {
	return tp.path.FullMatch()
}

func (tp templatePolicy) PopKey(key string) Policy {
	tp.path = tp.path.PopKey(key)
	return tp
}

func BuildPolicy(policyConfig *config.PolicyConfig, funcMap FuncMap) (Policy, error) {

	goTpl, goTplError := template.CreateTemplate("template", funcMap, policyConfig.Template)

	if goTplError != nil {
		return nil, goTplError
	}

	var unitTpls []Template

	for _, unitTemplate := range policyConfig.UnitTemplate {
		unitTpl, unitTplError := template.CreateTemplate("unit_template", funcMap, unitTemplate)

		if unitTplError != nil {
			return nil, unitTplError
		}
		unitTpls = append(unitTpls, unitTpl)
	}

	path := policy_path.Create(policyConfig.Path)

	return templatePolicy{path: path, tpl: goTpl, unitTpls: unitTpls}, nil
}
