package main

import (
	"fmt"
	"os"
	"path"

	"golang.org/x/exp/slices"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		fmt.Fprintf(os.Stderr, "usage: tfmod-list <directory>\n")
		os.Exit(1)
	}

	modules := collectModules(dir)
	slices.Sort(modules)
	uniqueModules := slices.Compact(modules)
	for _, v := range uniqueModules {
		fmt.Printf("%v\n", v)
	}
}

func collectModules(dir string) []string {
	var modules []string
	module, _ := tfconfig.LoadModule(dir)
	for _, v := range module.ModuleCalls {
		moduleSource := path.Join(dir, v.Source)
		modules = append(modules, moduleSource)

		// Collect modules in module
		modules = append(modules, collectModules(moduleSource)...)
	}

	return modules
}
