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
	"bytes"
	"github.com/synchronoss/dredger/components/document_leaf"
	"html"
	"html/template"
)

type FuncMap = template.FuncMap
type DocumentLeaf = document_leaf.DocumentLeaf

//
// Template
//
// Just wraps around html/template and makes it a little more sane

type Template interface {
	Execute(*DocumentLeaf) (string, error)
}

type goTemplate struct {
	goTpl *template.Template
}

func (tpl goTemplate) Execute(leaf *DocumentLeaf) (string, error) {
	var buf bytes.Buffer
	err := tpl.goTpl.Execute(&buf, leaf)
	if err != nil {
		return "", err
	}
	return html.UnescapeString(buf.String()), nil
}

func CreateTemplate(name string, funcMap FuncMap, body string) (Template, error) {
	tpl, tplError := template.New(name).Funcs(funcMap).Parse(body)

	if tplError != nil {
		return nil, tplError
	}

	return &goTemplate{goTpl: tpl}, nil
}
