package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func readImageByParse() (string, error) {
	f, err := ioutil.TempFile("", "pngpaste-*.png")
	if err != nil {
		return "", err
	}

	stderr := new(bytes.Buffer)
	cmd := exec.Command("pngpaste", f.Name())
	cmd.Stderr = stderr
	if err = cmd.Run(); err != nil {
		return "", fmt.Errorf("err: %s, status: %s", strings.TrimSpace(stderr.String()), strings.TrimSpace(err.Error()))
	}

	return f.Name(), nil
}
