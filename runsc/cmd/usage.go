// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"

	"github.com/google/subcommands"
	"gvisor.dev/gvisor/runsc/config"
	"gvisor.dev/gvisor/runsc/container"
	"gvisor.dev/gvisor/runsc/flag"
)

// Usage implements subcommands.Command for the "usage" command.
type Usage struct {
	full bool
}

// Name implements subcommands.Command.Name.
func (*Usage) Name() string {
	return "usage"
}

// Synopsis implements subcommands.Command.Synopsis.
func (*Usage) Synopsis() string {
	return "Usage shows application memory usage across various categories in bytes."
}

// Usage implements subcommands.Command.Usage.
func (*Usage) Usage() string {
	return `usage [flags] <container id> - print memory usages to standard output.`
}

// SetFlags implements subcommands.Command.SetFlags.
func (c *Usage) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.full, "full", false, "enumerate all usage by categories")
}

// Execute implements subcommands.Command.Execute.
func (c *Usage) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() < 1 {
		f.Usage()
		return subcommands.ExitUsageError
	}

	id := f.Arg(0)
	conf := args[0].(*config.Config)

	cont, err := container.Load(conf.RootDir, container.FullID{ContainerID: id}, container.LoadOpts{})
	if err != nil {
		Fatalf("loading container: %v", err)
	}

	if err := cont.Usage(c.full); err != nil {
		Fatalf("usage failed: %v", err)
	}

	return subcommands.ExitSuccess
}
