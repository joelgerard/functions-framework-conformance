// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This binary contains a valiation framework for functions frameworks.
package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	cmd             = flag.String("cmd", "", "command to run a Functions Framework server at localhost:8080")
	functionType    = flag.String("type", "http", "type of function to validate (must be 'http', 'cloudevent', or 'legacyevent'")
	validateMapping = flag.Bool("validate-mapping", true, "whether to validate mapping from legacy->cloud events and vice versa (as applicable)")
)

func runValidation() error {
	log.Printf("Validating %q for %s...", *cmd, *functionType)

	shutdown, err := start(*cmd)
	defer shutdown()

	if err != nil {
		return fmt.Errorf("unable to start server: %v", err)
	}

	if err := validate("http://localhost:8080", *functionType, *validateMapping); err != nil {
		return fmt.Errorf("Validation failure: %v", err)
	}

	log.Printf("All validation passed!")
	return nil
}

func main() {
	flag.Parse()
	if err := runValidation(); err != nil {
		log.Fatalf("%v", err)
	}
}
