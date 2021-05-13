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
package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	baseConfig := LoadDefaultConfig()

	if baseConfig == "" {
		t.Fatalf("LoadDefaultConfig returned empty")
	}

	c, err := unmarshalConfig(baseConfig)

	if err != nil {
		t.Fatal(err)
	}

	if c == nil {
		t.Fatalf("Unmarshaled config is nil")
	}
}
