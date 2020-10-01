package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"sync"
	"time"
)

var command = &cobra.Command{
	Use:   "cssh",
	Short: "ssh, but concurrent",
	Long:  `A simple tool to execute ssh commands against multiple hosts`,
	Run: func(cmd *cobra.Command, args []string) {
		ssh([]string{"host1", "host2"})
	},
}

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ssh(hosts []string) {
	// todo: allow this as input with reasonable default
	limit := 1

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
			fmt.Println("executing on host:", host)

			// todo: args
			cmd := exec.Command("ssh", host)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", host, err)
			}

			time.Sleep(time.Second * 1)

			// receive from channel to free up a spot
			<-limitChannel
		}(i)
	}
	wg.Wait()
}
