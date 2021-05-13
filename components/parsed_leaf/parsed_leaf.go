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
package parsed_leaf

//
// ParsedLeaf represents a POST-rendering document leaf.
//
// It also includes additional units which have to propegate
// up to the root.
//

type ParsedLeaf interface {
	GetLeaf() string
	GetUnits() []string
	GetValue() interface{}
	GetAllUnits() []string
	SetLeaf(string) ParsedLeaf
	SetValue(interface{}) ParsedLeaf
	AddUnits(...string) ParsedLeaf
	JoinUnits(ParsedLeaf) ParsedLeaf
}

type parsedLeaf struct {
	leaf  string
	units []string
	value interface{}
}

func (l parsedLeaf) GetValue() interface{} { return l.value }
func (l parsedLeaf) GetLeaf() string       { return l.leaf }
func (l parsedLeaf) GetUnits() []string    { return l.units }
func (l parsedLeaf) GetAllUnits() []string { return append(l.units, l.leaf) }

func (l parsedLeaf) SetLeaf(newLeaf string) ParsedLeaf {
	l.leaf = newLeaf
	return l
}

func (l parsedLeaf) SetValue(newValue interface{}) ParsedLeaf {
	l.value = newValue
	return l
}

func (l parsedLeaf) AddUnits(newUnits ...string) ParsedLeaf {
	l.units = append(l.units, newUnits...)
	return l
}

func (l parsedLeaf) JoinUnits(r ParsedLeaf) ParsedLeaf {
	return l.AddUnits(r.GetUnits()...)
}

func Create(leaf string) ParsedLeaf {
	return &parsedLeaf{leaf: leaf}
}
