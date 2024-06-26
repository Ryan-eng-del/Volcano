package main

import (
	"volcano.user_srv/cmd"
)


func main() {
	// Get command line arguments
	if err := cmd.CmdExecute(); err != nil {
		panic(err)
	}
}