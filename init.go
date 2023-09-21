package main

import (
	"os"

	"github.com/cedana/runc/libcontainer"
	_ "github.com/cedana/runc/libcontainer/nsenter"
)

func init() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		// This is the golang entry point for runc init, executed
		// before main() but after libcontainer/nsenter's nsexec().
		libcontainer.Init()
	}
}
