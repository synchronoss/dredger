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
package debug

import (
	"fmt"
	"os"
)

// Just a debug function which can be enabled via
// an env var DEBUG=on

var printDebug = os.Getenv("DEBUG") == "on"

func Debug(msg ...interface{}) {
	if printDebug {
		fmt.Println(msg...)
	}
}
