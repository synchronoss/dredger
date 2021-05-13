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
package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {

	template, templateErr := CreateTemplate("test", FuncMap{}, "Key = {{ $.Key }}")

	if templateErr != nil {
		t.Fatal(templateErr)
	}

	docLeaf := DocumentLeaf{
		Key: "foo",
	}

	output, err := template.Execute(&docLeaf)

	if err != nil {
		t.Fatal(err)
	}

	if output != "Key = foo" {
		t.Fatalf("Unexpected output, got %s", output)
	}
}
