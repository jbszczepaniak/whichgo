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

type GitRepo struct{url string}

func (gr GitRepo) GoVer() error{
	if err := gr.validate(); err != nil {
		return err
	}
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	cmd := exec.Command("git", "clone", gr.url, dir)
	cmd.Dir = dir
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("git", "ls-files")
	cmd.Dir = dir
	out, err := cmd.Output()
	br := bufio.NewReader(bytes.NewReader(out))
	for {
		read, err := br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		trimmed := strings.TrimSuffix(string(read), "\n")
		if strings.HasSuffix(trimmed, "go.mod") {
			fmt.Println(goVerFromMod(filepath.Join(dir, trimmed)))
		}
	}
	return nil
}

func (gr GitRepo) validate() error {
	cmd := exec.Command("git", "ls-remote", gr.url)
	cmd.Env = gitEnvNoPrompt
	err := cmd.Run()
	return err
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
