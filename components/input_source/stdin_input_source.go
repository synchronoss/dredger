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
package input_source

//
// stdinInputSource reads all documents from stdin and assumes YAML encoding
//

import (
	"github.com/synchronoss/dredger/components/encoding"
	"io/ioutil"
	"os"
)

type stdinInputSource struct{}

func (si stdinInputSource) ReadDocuments() ([]interface{}, error) {

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		return nil, err
	}

	asString := string(bytes)

	return encoding.DecodeDocumentList(asString)
}
