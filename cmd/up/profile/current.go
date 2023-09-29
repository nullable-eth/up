// Copyright 2022 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package profile

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/upbound/up/internal/profile"
	"github.com/upbound/up/internal/upbound"
)

type currentCmd struct{}

type output struct {
	Name    string           `json:"name"`
	Profile profile.Redacted `json:"profile"`
}

// Run executes the current command.
func (c *currentCmd) Run(ctx *kong.Context, upCtx *upbound.Context) error {
	name, prof, err := upCtx.Cfg.GetDefaultUpboundProfile()
	if err != nil {
		return err
	}

	redacted := profile.Redacted{Profile: prof}

	b, err := json.MarshalIndent(output{
		Name:    name,
		Profile: redacted,
	}, "", "    ")
	if err != nil {
		return err
	}
	fmt.Fprintln(ctx.Stdout, string(b))
	return nil
}
