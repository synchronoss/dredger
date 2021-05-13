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
package document_leaf

//
// Document leaf is the wrapper structure
// used to present data to the templating engine
//
// It consists of the "actual" values and the "substitute"
// values that would have been rendered had a policy not matched.
//
// As a result this is basically a Leaf and ParsedLeaf combined.
// But due to the way that go templating works they are presented
// as members of a single object

type DocumentLeaf struct {
	Key      string
	Doc      interface{}
	Sub      string
	Trail    string
	SubValue interface{}
	Value    interface{}
}
