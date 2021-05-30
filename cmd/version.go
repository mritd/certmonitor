package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionTpl = `
Name: certmonitor
Version: %s
Arch: %s
BuildTime: %s
CommitID: %s
`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long: `
Print version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionTpl, version, runtime.GOOS+"/"+runtime.GOARCH, buildTime, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
