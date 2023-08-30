package main

import (
	"github.com/Zelayan/dts/cmd/colletcor/command"
	"k8s.io/klog/v2"
	"os"
)

func main() {
	klog.InitFlags(nil)
	cmd := command.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
