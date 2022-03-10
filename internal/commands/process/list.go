// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package process

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/apache/skywalking-cli/internal/commands/interceptor"
	"github.com/apache/skywalking-cli/internal/flags"
	"github.com/apache/skywalking-cli/pkg/display"
	"github.com/apache/skywalking-cli/pkg/display/displayable"
	"github.com/apache/skywalking-cli/pkg/graphql/metadata"
)

var ListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   `list all monitored processes of the given id or name in service or instance`,
	UsageText: `This command lists all processes of the service-instance, via id or name in service or instance.

Examples:
1. List all processes by service name "provider":
$ swctl process ls --service-name provider

2. List all processes by service id "YnVzaW5lc3Mtem9uZTo6cHJvamVjdEM=.1":
$ swctl process ls --service-id YnVzaW5lc3Mtem9uZTo6cHJvamVjdEM=.1

3. List all processes by instance name "provider-01" and service name "provider":
$ swctl process ls --instance-name provider-01 --service-name provider

4. List all processes by instance id "cHJvdmlkZXI=.1_cHJvdmlkZXIx":
$ swctl process ls --service-id cHJvdmlkZXI=.1_cHJvdmlkZXIx`,
	Flags: flags.Flags(
		flags.ServiceFlags,
		flags.InstanceFlags,
	),
	Before: interceptor.BeforeChain(
		interceptor.ParseService(false),
		interceptor.ParseInstance(false),
	),
	Action: func(ctx *cli.Context) error {
		serviceID := ctx.String("service-id")
		instanceID := ctx.String("instance-id")
		if serviceID == "" && instanceID == "" {
			return fmt.Errorf("service or instance must provide one")
		}

		processes, err := metadata.Processes(ctx, serviceID, instanceID)
		if err != nil {
			return err
		}

		return display.Display(ctx, &displayable.Displayable{Data: processes})
	},
}