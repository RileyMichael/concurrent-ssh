package main

import (
	"fmt"
	"github.com/rileymichael/concurrent-ssh/pkg/concurrent"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var command = &cobra.Command{
	Use:   "cssh",
	Short: "ssh, but concurrent",
	Long: `A simple tool to execute ssh commands against multiple hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		targets, limit := processInput(cmd)
		commands := concurrent.BuildCommands("ssh", targets, args)
		concurrent.ExecuteCommands(commands, limit)
	},
}

func init() {
	command.Flags().StringP("file", "f", "", "file containing target hosts")
	command.Flags().StringSliceP("targets", "t", nil, "comma separated target hosts")
	command.Flags().IntP("limit", "l", 25, "concurrency limit (default 25)")
}

func processInput(cmd *cobra.Command) ([]string, int) {
	file, _ := cmd.Flags().GetString("file")
	targets, _ := cmd.Flags().GetStringSlice("targets")
	limit, _ := cmd.Flags().GetInt("limit")

	if file != "" && len(targets) > 0 {
		fmt.Println("Both file and targets input supplied; provide only 1.")
		os.Exit(1)
	} else if file == "" && len(targets) == 0 {
		fmt.Println("Neither file or targets input supplied; provide 1.")
		os.Exit(1)
	}

	if file != "" {
		data, err := ioutil.ReadFile(file)

		if err != nil {
			fmt.Println("Error reading hosts file")
			os.Exit(1)
		}

		// trim last newline if exists & split on newlines
		targets = strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	}

	return targets, limit
}

func main() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
