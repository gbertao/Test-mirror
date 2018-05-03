// Copyright (C) 2018 The Android Open Source Project
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

package selinux

// This file contains "se_cil_compat_map" module type used to build and install
// sepolicy backwards compatibility mapping files.

import (
	"android/soong/android"
	"fmt"
	"io"
)

var (
	pctx = android.NewPackageContext("android/soong/selinux")
)

func init() {
	android.RegisterModuleType("se_cil_compat_map", cilCompatMapFactory)
	pctx.Import("android/soong/common")
}

func cilCompatMapFactory() android.Module {
	c := &cilCompatMap{}
	c.AddProperties(&c.properties)
	android.InitAndroidModule(c)
	return c
}

type cilCompatMapProperties struct {
	// list of source (.cil) files used to build an sepolicy compatibility mapping
	// file. srcs may reference the outputs of other modules that produce source
	// files like genrule or filegroup using the syntax ":module". srcs has to be
	// non-empty.
	Srcs []string
}

type cilCompatMap struct {
	android.ModuleBase
	properties cilCompatMapProperties
	// (.intermediate) module output path as installation source.
	installSource android.OptionalPath
}

func (c *cilCompatMap) GenerateAndroidBuildActions(ctx android.ModuleContext) {
	srcFiles := ctx.ExpandSources(c.properties.Srcs, nil)
	for _, src := range srcFiles {
		if src.Ext() != ".cil" {
			ctx.PropertyErrorf("srcs", "%s has to be a .cil file.", src.String())
		}
	}

	out := android.PathForModuleGen(ctx, c.Name())
	ctx.Build(pctx, android.BuildParams{
		Rule:   android.Cat,
		Output: out,
		Inputs: srcFiles,
	})
	c.installSource = android.OptionalPathForPath(out)
}

func (c *cilCompatMap) DepsMutator(ctx android.BottomUpMutatorContext) {
	android.ExtractSourcesDeps(ctx, c.properties.Srcs)
}

func (c *cilCompatMap) AndroidMk() android.AndroidMkData {
	ret := android.AndroidMkData{
		OutputFile: c.installSource,
		Class:      "ETC",
	}
	ret.Extra = append(ret.Extra, func(w io.Writer, outputFile android.Path) {
		fmt.Fprintln(w, "LOCAL_MODULE_PATH := $(TARGET_OUT)/etc/selinux/mapping")
	})
	return ret
}
