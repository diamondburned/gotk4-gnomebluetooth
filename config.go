package main

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module = "github.com/diamondburned/gotk4/pkg"
	thisModule  = "github.com/diamondburned/gotk4-gnomebluetooth/pkg"
)

var packages = []gendata.Package{
	{PkgName: "gnome-bluetooth-1.0", Namespaces: nil},
}

var pkgGenerated = []string{
	"gnomebluetooth",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var preprocessors = []types.Preprocessor{}

var filters = []types.FilterMatcher{}
