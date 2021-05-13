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

package main

import (
	"flag"
	"fmt"
	"github.com/synchronoss/dredger/components/config"
	"github.com/synchronoss/dredger/components/converter_logic"
	"github.com/synchronoss/dredger/components/input_source"
	"github.com/synchronoss/dredger/components/leaf"
	"github.com/synchronoss/dredger/components/output_target"
	"github.com/synchronoss/dredger/components/parsed_leaf"
	"github.com/synchronoss/dredger/components/policy_set"
	"github.com/synchronoss/dredger/components/template_function_map"
	"os"
)

const (
	version = "0.4.0"
)

var (
	dumpConfig  bool
	versionFlag bool
	configFile  string
	outputDir   string
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	flag.BoolVar(&dumpConfig, "dumpconfig", false, "Dump the default config")
	flag.BoolVar(&versionFlag, "version", false, "Display the version and exit")
	flag.StringVar(&configFile, "config", "", "Specify a new config file")
	flag.StringVar(&outputDir, "outputdir", "", "Split units into files and write to dir")
	flag.Parse()

	if dumpConfig {
		baseConfig := config.LoadDefaultConfig()
		fmt.Println(baseConfig)
		os.Exit(0)
	}

	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	args := flag.Args()

	var mode string = "stdin"
	var extraArgs []string

	if len(args) > 0 {
		mode = args[0]
		extraArgs = args[1:]
	}

	inputSource, inputSourceErr := input_source.CreateInputSource(mode, extraArgs)
	panicOnErr(inputSourceErr)

	documents, decodeErr := inputSource.ReadDocuments()
	panicOnErr(decodeErr)

	config, configErr := config.LoadConfig(configFile)
	panicOnErr(configErr)

	funcMap := template_function_map.CreateFunctionMap()

	rootPolicySet, rootPolicyErr := policy_set.BuildPolicies(config, funcMap)
	panicOnErr(rootPolicyErr)

	outputTarget := output_target.CreateOutputTarget(outputDir)

	converterLogic := converter_logic.CreateDefaultConverterLogic()

	for _, staticUnit := range config.StaticUnits {
		staticLeaf := parsed_leaf.Create(staticUnit)
		outputTarget.WriteLeaf(staticLeaf)
	}

	for _, doc := range documents {
		l := leaf.CreateLeaf(doc, converterLogic, rootPolicySet)

		parsedLeaf, parseErr := l.Process()
		panicOnErr(parseErr)

		outputErr := outputTarget.WriteLeaf(parsedLeaf)
		panicOnErr(outputErr)
	}

}
