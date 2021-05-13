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
	"errors"
	"github.com/synchronoss/dredger/components/debug"
	"io/ioutil"
)

//
// An output target to write each leaf/unit file to it's own file
//

type directoryOutputTarget struct {
	directory string
}

func (dt directoryOutputTarget) determinFileName(unitPath string) string {
	return dt.directory + "/" + unitPath + ".tf"
}

func (dt directoryOutputTarget) writeWorldReadableFile(fileName string, data string) error {
	return ioutil.WriteFile(fileName, []byte(data), 0644)
}

func (dt directoryOutputTarget) writeUnitFile(unitPath string, data string) error {
	fileName := dt.determinFileName(unitPath)
	debug.Debug("writing to ", fileName)
	return dt.writeWorldReadableFile(fileName, data)
}

func (do directoryOutputTarget) WriteLeaf(p ParsedLeaf) error {
	for _, unit := range p.GetAllUnits() {
		unitPath := DetermineUnitPath(unit)

		if unitPath == "" {
			return errors.New("Failed to determine unit path")
		}

		writeErr := do.writeUnitFile(unitPath, unit)

		if writeErr != nil {
			return writeErr
		}
	}

	return nil
}
