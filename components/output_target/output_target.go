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
package output_target

import (
	"github.com/synchronoss/dredger/components/parsed_leaf"
)

//
// OutputTarget is any destination for converted documents
//

type ParsedLeaf = parsed_leaf.ParsedLeaf

type OutputTarget interface {
	Init() error
	WriteLeaf(ParsedLeaf) error
}

func CreateOutputTarget(outputDir string) OutputTarget {
	if outputDir == "" {
		return stdoutOutputTarget{}
	} else {
		return directoryOutputTarget{directory: outputDir}
	}
}
