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
package converter_logic

import (
	"fmt"
	"github.com/synchronoss/dredger/components/map_utils"
	"github.com/synchronoss/dredger/components/terraform_format"
)

//
// converter logic for terraform "map" format e.g.
//
// section = {
//   "key" = "value"
//   "foo" = [ { "bar":22 }, { "baz":420 } ]
// }
//

type mapConverterLogic struct{}

func (dl mapConverterLogic) FormatKey(k string) string {
	return terraform_format.Quote(terraform_format.SnakeCase(k))
}

func (dl mapConverterLogic) FormatMap(k string, m map[string]string) string {

	var content = ""

	// Itterate through the keys in sorted order (for consistency)
	for _, k := range map_utils.SortKeysString(m) {

		// The sub-value already contains the key so just append
		// to the content with a newline.
		content += m[k] + "\n"
	}

	// Indent the whole content.
	indentContent := terraform_format.Indent(content)

	// Build the braced leaf-string
	key := terraform_format.AssignKey(k)
	return fmt.Sprintf("%s {\n%s\n}", key, indentContent)
}

func (dl mapConverterLogic) FormatArray(k string, a []string) string {
	var content = ""

	for _, v := range a {
		content += v + ",\n"
	}

	indentContent := terraform_format.Indent(content)

	key := terraform_format.AssignKey(k)
	return fmt.Sprintf("%s [\n%s\n]", key, indentContent)
}

func (dl mapConverterLogic) FormatString(k string, v string) string {
	key := terraform_format.AssignKey(k)
	s := terraform_format.FormatStringValue(v)
	return fmt.Sprintf("%s %s", key, s)
}

func (dl mapConverterLogic) FormatInt(k string, v int) string {
	key := terraform_format.AssignKey(k)
	return fmt.Sprintf("%s %d", key, v)
}

func (dl mapConverterLogic) FormatFloat(k string, v float64) string {
	key := terraform_format.AssignKey(k)
	return fmt.Sprintf("%s %f", key, v)
}

func (dl mapConverterLogic) FormatBool(k string, v bool) string {
	key := terraform_format.AssignKey(k)
	return fmt.Sprintf("%s %t", key, v)
}

func (dl mapConverterLogic) FormatNill(k string) string {
	key := terraform_format.AssignKey(k)
	return fmt.Sprintf("%s null", key)
}

func CreateMapConverterLogic() ConverterLogic {
	return mapConverterLogic{}
}
