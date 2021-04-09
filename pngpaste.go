package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func readImageByParse() (string, error) {
	pngBin, err := getPngBin()
	if err != nil {
		return "", err
	}

	f, err := ioutil.TempFile("", "pngpaste-*.png")
	if err != nil {
		return "", err
	}

	stderr := new(bytes.Buffer)
	cmd := exec.Command(pngBin, f.Name())
	cmd.Stderr = stderr
	if err = cmd.Run(); err != nil {
		return "", fmt.Errorf("err: %s, status: %s", strings.TrimSpace(stderr.String()), strings.TrimSpace(err.Error()))
	}

	return f.Name(), nil
}

func getPngBin() (string, error) {
	p, err := exec.LookPath("pngpaste")
	if err != nil {
		cmd := exec.Command("brew", "install", "pngpaste")
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return "", err
		}
		p, err = exec.LookPath("pngpaste")
	}

	return p, err
}
