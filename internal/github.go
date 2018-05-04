package internal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func getSHA(repo, token, path string) string {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path), nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var sha struct {
		SHA string `json:"sha"`
	}
	err = json.Unmarshal(body, &sha)
	if err != nil {
		return ""
	}

	return sha.SHA
}

func Upload(repo, token, path string, content []byte) error {
	sha := getSHA(repo, token, path)
	var data io.Reader
	if sha == "" {
		data = bytes.NewReader([]byte(fmt.Sprintf(`{
  "message": "upload %s",
  "content": "%s"
}`, path, base64.StdEncoding.EncodeToString(content))))
	} else {
		data = bytes.NewReader([]byte(fmt.Sprintf(`{
  "message": "upload %s",
  "content": "%s",
  "sha": "%s"
}`, path, base64.StdEncoding.EncodeToString(content), sha)))
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://api.github.com/repos/%s/contents/%s", repo, path), data)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf(string(body))
	}

	return nil
}
