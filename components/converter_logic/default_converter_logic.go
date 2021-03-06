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
// This creates what we consider to be the "default" dredger
// converter logic. In this case convert arrays to maps and
// everything else to block format. e.g.
//
// section {
//   foo = "value"
//   bar = [1,2,3]
// }
//

func CreateDefaultConverterLogic() ConverterLogic {
	mapConverter := mapConverterLogic{}
	blockConverter := blockConverterLogic{}

	return joinedConverterLogic{
		mapConverter:    blockConverter,
		arrayConverter:  mapConverter,
		scalarConverter: blockConverter,
	}
}
