// Copyright (C) 2014 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("earthquake", "0.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"orchestrator": orchestratorCommandFactory,
		"guestagent":   guestAgentCommandFactory,
		"searchtools":  searchToolsCommandFactory,

		"init": initCommandFactory,
		"run":  runCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Printf("failed to execute command: %s\n", err)
	}

	os.Exit(exitStatus)
}
