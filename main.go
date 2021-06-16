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
	version = "0.4.3"
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

	// Parse the command line flags
	flag.BoolVar(&dumpConfig, "dumpconfig", false, "Dump the default config")
	flag.BoolVar(&versionFlag, "version", false, "Display the version and exit")
	flag.StringVar(&configFile, "config", "", "Specify a new config file")
	flag.StringVar(&outputDir, "outputdir", "", "Split units into files and write to dir")
	flag.Parse()

	// If dumpconfig flag was passed just write the config and exit
	if dumpConfig {
		baseConfig := config.LoadDefaultConfig()
		fmt.Println(baseConfig)
		os.Exit(0)
	}

	// If version flag was passed just write the version and exit
	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	args := flag.Args()

	// Determine the input mode and seperate out the helm arguments
	var mode string = "stdin"
	var extraArgs []string

	if len(args) > 0 {
		mode = args[0]
		extraArgs = args[1:]
	}

	// Create input source
	inputSource, inputSourceErr := input_source.CreateInputSource(mode, extraArgs)
	panicOnErr(inputSourceErr)

	// Read documents from input source
	documents, decodeErr := inputSource.ReadDocuments()
	panicOnErr(decodeErr)

	// Load the config
	config, configErr := config.LoadConfig(configFile)
	panicOnErr(configErr)

	// Generate template function map
	funcMap := template_function_map.CreateFunctionMap()

	// Build the initial policy set
	rootPolicySet, rootPolicyErr := policy_set.BuildPolicies(config, funcMap)
	panicOnErr(rootPolicyErr)

	// Create the output target
	outputTarget := output_target.CreateOutputTarget(outputDir)

	// Initialize the output target
	outputTargetInitErr := outputTarget.Init()
	panicOnErr(outputTargetInitErr)

	// We need a default logic converter
	converterLogic := converter_logic.CreateDefaultConverterLogic()

	// If we reach this point then we always output the static units
	for _, staticUnit := range config.StaticUnits {
		staticLeaf := parsed_leaf.Create(staticUnit)
		outputTarget.WriteLeaf(staticLeaf)
	}

	// Loop through the documents
	for _, doc := range documents {
		// Convert to terraform
		l := leaf.CreateLeaf(doc, converterLogic, rootPolicySet)
		parsedLeaf, parseErr := l.Process()
		panicOnErr(parseErr)

		// Write to the output target
		outputErr := outputTarget.WriteLeaf(parsedLeaf)
		panicOnErr(outputErr)
	}

}
