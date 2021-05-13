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
// An interface for the converter logic.
//
// converter logic classes must specify what to do for each encountered
// base type.
//
// These functions act upon branches where the sub-branches have already
// been converted, so each container member has a string sub-type.

type ConverterLogic interface {
	FormatKey(string) string

	FormatMap(string, map[string]string) string
	FormatArray(string, []string) string
	FormatString(string, string) string
	FormatInt(string, int) string
	FormatFloat(string, float64) string
	FormatBool(string, bool) string
	FormatNill(string) string
}
