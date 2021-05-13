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
package leaf

import (
	"fmt"
	"github.com/synchronoss/dredger/components/converter_logic"
	"github.com/synchronoss/dredger/components/debug"
	"github.com/synchronoss/dredger/components/document_leaf"
	"github.com/synchronoss/dredger/components/parsed_leaf"
	"github.com/synchronoss/dredger/components/policy_set"
)

type Leaf interface {
	Process() (parsed_leaf.ParsedLeaf, error)
}

func CreateLeaf(doc interface{}, logic converter_logic.ConverterLogic, policies policy_set.PolicySet) Leaf {
	return leaf{
		key:      "",
		trail:    "",
		doc:      doc,
		value:    doc,
		logic:    logic,
		policies: policies,
	}
}

func CreateConverterLeaf(doc interface{}, key string, logic converter_logic.ConverterLogic) Leaf {
	return leaf{
		key:      key,
		trail:    key,
		doc:      doc,
		value:    doc,
		logic:    logic,
		policies: policy_set.CreateEmptyPolicySet(),
	}
}

type leaf struct {
	logic    converter_logic.ConverterLogic
	policies policy_set.PolicySet
	trail    string
	key      string
	value    interface{}
	doc      interface{}
}

func (l leaf) Process() (parsed_leaf.ParsedLeaf, error) {
	subLeaf, subErr := l.determineSub()

	if subErr != nil {
		return nil, subErr
	}

	if l.policies.Match() {
		debug.Debug("!! policy match at", l.trail)

		docLeaf := document_leaf.DocumentLeaf{
			Key:      l.key,
			Trail:    l.trail,
			Value:    l.value,
			Sub:      subLeaf.GetLeaf(),
			Doc:      l.doc,
			SubValue: subLeaf.GetValue(),
		}

		policyLeaf, policyErr := l.policies.Execute(&docLeaf)

		if policyErr != nil {
			return nil, policyErr
		}

		return policyLeaf.JoinUnits(subLeaf), nil
	} else {
		return subLeaf, nil
	}
}

func (l leaf) determineSub() (parsed_leaf.ParsedLeaf, error) {

	v := l.value

	if v == nil {
		return l.convertNill()
	}

	switch v.(type) {
	case map[string]interface{}:
		return l.convertMap()
	case []interface{}:
		return l.convertArray()
	case string:
		return l.convertString()
	case int:
		return l.convertInt()
	case float64:
		return l.convertFloat()
	case bool:
		return l.convertBool()
	default:
		return nil, fmt.Errorf("Unknown type produced in YAML: %T", v)
	}

}

func (l leaf) convertNill() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting nill at", l.trail)
	stringValue := l.logic.FormatNill(l.key)
	return parsed_leaf.Create(stringValue), nil
}

func (l leaf) convertMap() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting map at", l.trail)
	asMap := l.value.(map[string]interface{})

	// Detremine the "sub" values first.
	sub := map[string]string{}

	// An empty parsed_leaf.ParsedLeaf is created
	// to merge the sub values with
	var parsed = parsed_leaf.Create("")
	for k, v := range asMap {

		subPolicySet := l.policies.Reduce(k)

		formatKey := l.logic.FormatKey(k)

		subLeaf := leaf{
			key:      formatKey,
			value:    v,
			logic:    l.logic,
			policies: subPolicySet,
			doc:      l.doc,
			trail:    l.trail + "." + formatKey,
		}

		subOut, err := subLeaf.Process()

		if err != nil {
			return subOut, err
		}

		// Join the sub-parsed-leaf's units to
		// our sub-value.
		parsed = parsed.JoinUnits(subOut)

		subText := subOut.GetLeaf()

		if subText == "" {
			continue
		}

		// Add to the sub-map
		sub[k] = subText
	}

	formatted := l.logic.FormatMap(l.key, sub)

	return parsed.SetLeaf(formatted).SetValue(sub), nil
}

func (l leaf) convertArray() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting array at", l.trail)
	asArray := l.value.([]interface{})

	// Determine the "sub" value

	var sub []string
	var parsed = parsed_leaf.Create("")

	// Iterate through the items and convert
	for i, v := range asArray {
		k := fmt.Sprintf("%d", i)

		// Reduce the policies for this key
		subPolicySet := l.policies.Reduce(k)

		// Create a sub-leaf
		subLeaf := leaf{
			key:      "",
			value:    v,
			logic:    l.logic,
			policies: subPolicySet,
			doc:      l.doc,
			trail:    l.trail + "." + k,
		}

		subOut, err := subLeaf.Process()

		if err != nil {
			return nil, err
		}

		// Add the sub-units to the return leaf
		parsed = parsed.JoinUnits(subOut)

		subText := subOut.GetLeaf()

		if subText == "" {
			continue
		}

		sub = append(sub, subText)

	}

	formatted := l.logic.FormatArray(l.key, sub)

	return parsed.SetLeaf(formatted).SetValue(sub), nil
}

//
// These methods just refer to the converter logic but wrap in leaves.
//

func (l leaf) convertString() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting string at", l.trail)
	stringValue := l.logic.FormatString(l.key, l.value.(string))
	return parsed_leaf.Create(stringValue), nil
}

func (l leaf) convertInt() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting int at", l.trail)
	stringValue := l.logic.FormatInt(l.key, l.value.(int))
	return parsed_leaf.Create(stringValue), nil
}

func (l leaf) convertFloat() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting float at", l.trail)
	stringValue := l.logic.FormatFloat(l.key, l.value.(float64))
	return parsed_leaf.Create(stringValue), nil
}

func (l leaf) convertBool() (parsed_leaf.ParsedLeaf, error) {
	debug.Debug("converting bool at", l.trail)
	stringValue := l.logic.FormatBool(l.key, l.value.(bool))
	return parsed_leaf.Create(stringValue), nil
}
