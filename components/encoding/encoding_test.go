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
	"testing"
)

const testDoc = `
---
one: 1
---
two: 2
`

func TestEncoding(t *testing.T) {

	documents, err := DecodeDocumentList(testDoc)

	if err != nil {
		t.Fatal(err)
	}

	if 2 != len(documents) {
		t.Fatalf("Wrong number of documents, expected 2 got %d", len(documents))
	}

}
