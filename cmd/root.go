package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var command = &cobra.Command{
	Use:   "cssh",
	Short: "ssh, but concurrent",
	Long: `A simple tool to execute ssh commands against multiple hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo fella")
	},
}

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
