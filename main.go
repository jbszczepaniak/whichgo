package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
type sourceType string
const (
	git sourceType = "git repository"
	dir sourceType = "directory"
	mod sourceType = "go module"
)

var gitEnvNoPrompt = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")

// what is the type of given source?
func whatType(source string) sourceType {
	cmd := exec.Command("git", "ls-remote", source)
	cmd.Env = gitEnvNoPrompt
	err := cmd.Run()
	if err == nil {
		return git
	}
	panic("not a git repo, or do not have permissions")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("source is not provided")
	}
	source := os.Args[1]

	switch whatType(source) {
	case git:
		handleGit(source)
	case dir:
	case mod:
	}


}

func handleGit(source string) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	cmd := exec.Command("git", "clone", source, dir)
	cmd.Dir = dir
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("git", "ls-files", )
	cmd.Dir = dir
	out, err := cmd.Output()
	br := bufio.NewReader(bytes.NewReader(out))
	for {
		read, err := br.ReadBytes('\n')
		if err != nil{
			if err == io.EOF {
				break
			}
		}
		trimmed := strings.TrimSuffix(string(read), "\n")
		if strings.HasSuffix(trimmed, "go.mod") {
			fmt.Println(goVerFromMod(filepath.Join(dir, trimmed)))
		}
	}
}

func goVerFromMod(file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "go") {
			return line
		}
	}
	return "NOT FOUND"
}
