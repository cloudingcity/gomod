package gomod

import (
	"golang.org/x/mod/modfile"
)

type Module struct {
	Path    string
	Version string
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
}

type Require struct {
	Path     string
	Version  string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}

func Parse(data []byte) (*GoMod, error) {
	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, err
	}

	var m GoMod

	if mod.Module != nil {
		m.Module = Module{
			Path:    mod.Module.Mod.Path,
			Version: mod.Module.Mod.Version,
		}
	}

	if mod.Go != nil {
		m.Go = mod.Go.Version
	}

	for _, r := range mod.Require {
		m.Require = append(m.Require, Require{
			Path:     r.Mod.Path,
			Version:  r.Mod.Version,
			Indirect: r.Indirect,
		})
	}

	for _, x := range mod.Exclude {
		m.Exclude = append(m.Exclude, Module{
			Path:    x.Mod.Path,
			Version: x.Mod.Version,
		})
	}

	for _, r := range mod.Replace {
		m.Replace = append(m.Replace, Replace{
			Old: Module{Path: r.Old.Path, Version: r.Old.Version},
			New: Module{Path: r.New.Path, Version: r.New.Version},
		})
	}

	return &m, nil
}
