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

import "fmt"

//
// InputSource represents any source that can produce document lists
// for conversion.
//

type InputSource interface {
	ReadDocuments() ([]interface{}, error)
}

func CreateInputSource(mode string, args []string) (InputSource, error) {
	switch mode {
	case "stdin":
		return stdinInputSource{}, nil
	case "helm":
		return helmInputSource{ args:args }, nil
	default:
		return nil, fmt.Errorf("unknown input source: %s", mode)
	}
}
