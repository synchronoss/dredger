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

//
// Here the basic config structure is defined along with
// the function to unmarshal it from yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	StaticUnits []string       `yaml:"static_units"`
	Policies    []PolicyConfig `yaml:"policies"`
}

type PolicyConfig struct {
	Path         string   `yaml:"path"`
	Template     string   `yaml:"template"`
	UnitTemplate []string `yaml:"unit_template"`
}

func LoadConfig(configFile string) (*Config, error) {
	if configFile == "" {
		return loadDefaultConfig()
	} else {
		return loadConfigPath(configFile)
	}
}

func loadDefaultConfig() (*Config, error) {
	return unmarshalConfig(baseConfig)
}

func unmarshalConfig(configStr string) (*Config, error) {
	config := Config{}
	yamlErr := yaml.Unmarshal([]byte(configStr), &config)

	return &config, yamlErr
}

func loadConfigPath(configFile string) (*Config, error) {
	bytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	return unmarshalConfig(string(bytes))
}
