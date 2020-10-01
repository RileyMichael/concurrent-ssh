package concurrent

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type cmd struct {
	command *exec.Cmd
	host    string
}

func ExecuteCommands(commands []cmd, limit int) {
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
			if output, err := commands[i].command.Output(); err != nil {
				fmt.Printf("%s\nerror: %s\n", hostPrefix(commands[i].host, true), err)
			} else {
				fmt.Printf("%s\n%s", hostPrefix(commands[i].host, false), output)
			}

			// receive from channel to free up a spot
			<-limitChannel
		}(i)
	}
	wg.Wait()
}

func BuildCommands(command string, hosts []string, args []string) []cmd {
	path, err := exec.LookPath(command)

	if err != nil {
		fmt.Println("Error: unable to find executable", command)
		os.Exit(1)
	}

	commands := make([]cmd, len(hosts))
	for i, host := range hosts {
		commandArgs := append([]string{host}, args...)
		command := exec.Command(path, commandArgs...)
		commands[i] = cmd{command: command, host: host}
	}

	return commands
}

func hostPrefix(host string, error bool) string {
	reset := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"

	if error {
		return fmt.Sprintf("%s[%s]%s", red, host, reset)
	} else {
		return fmt.Sprintf("%s[%s]%s", green, host, reset)
	}
}
