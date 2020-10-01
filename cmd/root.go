package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var command = &cobra.Command{
	Use:   "cssh",
	Short: "ssh, but concurrent",
	Long:  `A simple tool to execute ssh commands against multiple hosts`,
	RunE: func(cmd *cobra.Command, args []string) error {
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
				return errors.New("error reading hosts file")
			}

			// trim last newline if exists & split on newlines
			targets = strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
		}

		ssh(targets, args, limit)
		return nil
	},
}

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	command.Flags().StringP("file", "f", "", "file containing target hosts")
	command.Flags().StringSliceP("targets", "t", nil, "comma separated target hosts")
	command.Flags().IntP("limit", "l", 25, "concurrency limit (default 25)")
}

func ssh(hosts []string, args []string, limit int) {
	// buffered channel used to limit concurrency
	limitChannel := make(chan struct{}, limit)

	var wg sync.WaitGroup

	for i := 0; i < len(hosts); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// this will block until it can send to the channel / spot opens
			limitChannel <- struct{}{}

			// execute command
			host := hosts[i]

			commandArgs := append([]string{host}, args...)
			cmd := exec.Command("ssh", commandArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", host, err)
			}

			// receive from channel to free up a spot
			<-limitChannel
		}(i)
	}
	wg.Wait()
}
