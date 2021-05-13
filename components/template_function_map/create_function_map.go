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
package template_function_map

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/synchronoss/dredger/components/converter_logic"
	"github.com/synchronoss/dredger/components/leaf"
	"github.com/synchronoss/dredger/components/policy_set"
	"github.com/synchronoss/dredger/components/terraform_format"
)

//
// Creates the in-built functions which we extend go-templates with.
//
// These include things such as toOctal.
//

type BoundConverter = func(interface{}, string) (string, error)

func bindLogicConverter(logic converter_logic.ConverterLogic) BoundConverter {

	return func(v interface{}, k string) (string, error) {
		l := leaf.CreateConverterLeaf(v, k, logic)

		p, err := l.Process()
		if err != nil {
			return "", err
		}

		return p.GetLeaf(), nil
	}
}

func CreateFunctionMap() *policy_set.FuncMap {
	funcMap := sprig.FuncMap()

	blockLogic := converter_logic.CreateBlockConverterLogic()
	funcMap["toTfBlock"] = bindLogicConverter(blockLogic)

	mapLogic := converter_logic.CreateMapConverterLogic()
	funcMap["toTfMap"] = bindLogicConverter(mapLogic)

	funcMap["toOctal"] = func(v int) string {
		return fmt.Sprintf("%#o", v)
	}

	funcMap["stripTfVars"] = terraform_format.StripVars
	funcMap["formatTfString"] = terraform_format.FormatStringValue

	return &funcMap
}
