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

import (
	"github.com/synchronoss/dredger/components/debug"
	"github.com/synchronoss/dredger/components/encoding"
	"os/exec"
	"bytes"
	"errors"
)

type helmInputSource struct {
	args []string
}

func (hi helmInputSource) ReadDocuments() ([]interface{}, error) {

	var docs []interface{}

	staticArgs := []string{"template", "${var.name}", "--namespace", "${var.namespace}"}

	args := append(staticArgs, hi.args...)

	path, err := exec.LookPath("helm")

	if err != nil {
		return docs, err
	}

	cmdArgs := append([]string{"helm"}, args...)

	debug.Debug("running helm with args", cmdArgs)

	cmd := exec.Cmd{
		Path: path,
		Args: cmdArgs,
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	runErr := cmd.Run()

	if runErr != nil {
		return docs, errors.New(string(stderr.Bytes()))
	}

	return encoding.DecodeDocumentList(string(stdout.Bytes()))
}
