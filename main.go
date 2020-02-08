package main

import (
	"fmt"
	"os"
)


var gitEnvNoPrompt = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")

func main() {
	if len(os.Args) < 2 {
		fmt.Println("source is not provided")
	}
	source := os.Args[1]

	// it is cheaper to call 1 endpoint that cloning whole repo,
	// so probably correct order would be
	// 0. directory
	// 1. go modules
	// 2. git repo (but it was implemented first...)
	fmt.Println(GitRepo{source}.GoVer())

}

// Source represents "a thing" for which it will be figured out, what go version is used
// or go version*s* *are* used.
type Source interface {
	// GoVer returns versions used in given thing
	GoVer(string)
}
