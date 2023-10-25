package main

import (
	"fmt"
	"github.com/Zelayan/dts/cmd/collector/command"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"os"
)

var (
	GitHash   string
	BuildTime string
	GoVersion string
)

func Version() *cobra.Command {

	command := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Git Commit Hash: %s \n", GitHash)
			fmt.Printf("Build TimeStamp: %s \n", BuildTime)
			fmt.Printf("GoLang Version: %s \n", GoVersion)
		},
	}
	return command
}

func main() {
	klog.InitFlags(nil)
	cmd := command.NewServerCommand()
	cmd.AddCommand(Version())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
