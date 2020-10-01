package concurrent

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func ExecuteCommands(commands []*exec.Cmd, limit int) {
	// buffered channel used to limit concurrency
	limitChannel := make(chan struct{}, limit)

	var wg sync.WaitGroup

	for i := 0; i < len(commands); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// this will block until it can send to the channel / spot opens
			limitChannel <- struct{}{}

			// execute command
			if err := commands[i].Run(); err != nil {
				fmt.Println("Error:", err)
			}

			// receive from channel to free up a spot
			<-limitChannel
		}(i)
	}
	wg.Wait()
}

func BuildCommands(command string, hosts []string, args []string) []*exec.Cmd {
	path, err := exec.LookPath(command)

	if err != nil {
		fmt.Println("Error: unable to find executable", command)
		os.Exit(1)
	}

	commands := make([]*exec.Cmd, len(hosts))
	for i, host := range hosts {
		commandArgs := append([]string{host}, args...)
		cmd := exec.Command(path, commandArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		commands[i] = cmd
	}

	return commands
}
