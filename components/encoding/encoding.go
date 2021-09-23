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
package encoding

import (
	"github.com/synchronoss/dredger/components/debug"
	"gopkg.in/yaml.v3"
	s "strings"
)

//
// Just a function to decode multi-document YAML
// into an array of interfaces.
//

func DecodeDocumentList(input string) ([]interface{}, error) {
	var ret = []interface{}{}

	// Split the documents by the yaml seperator
	documentStrings := s.Split(input, "\n---\n")

	for _, documentString := range documentStrings {

		// Skip if the document is empty
		if s.Trim(documentString, " \n\t") == "" {
			debug.Debug("skipping empty document")
			continue
		}

		documentStruct := map[string]interface{}{}

		yamlErr := yaml.Unmarshal([]byte(documentString), &documentStruct)
		if yamlErr != nil {
			return ret, yamlErr
		}

		if len(documentStruct) == 0 {
			continue
		}

		ret = append(ret, documentStruct)
	}

	return ret, nil
}
