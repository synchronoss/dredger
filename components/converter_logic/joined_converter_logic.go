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

//
// Joined converter logic is an interface coupler allowing different
// logic to be used for maps arrays and scalars.
//

type joinedConverterLogic struct {
	mapConverter    ConverterLogic
	arrayConverter  ConverterLogic
	scalarConverter ConverterLogic
}

func (dl joinedConverterLogic) FormatKey(k string) string {
	return dl.mapConverter.FormatKey(k)
}

func (dl joinedConverterLogic) FormatMap(k string, m map[string]string) string {
	return dl.mapConverter.FormatMap(k, m)
}

func (dl joinedConverterLogic) FormatArray(k string, a []string) string {
	return dl.arrayConverter.FormatArray(k, a)
}

func (dl joinedConverterLogic) FormatString(k string, v string) string {
	return dl.scalarConverter.FormatString(k, v)
}

func (dl joinedConverterLogic) FormatInt(k string, v int) string {
	return dl.scalarConverter.FormatInt(k, v)
}

func (dl joinedConverterLogic) FormatFloat(k string, v float64) string {
	return dl.scalarConverter.FormatFloat(k, v)
}

func (dl joinedConverterLogic) FormatBool(k string, v bool) string {
	return dl.scalarConverter.FormatBool(k, v)
}

func (dl joinedConverterLogic) FormatNill(k string) string {
	return dl.scalarConverter.FormatNill(k)
}
