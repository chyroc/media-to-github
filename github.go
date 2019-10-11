package main

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

func upload(repo, token, path string, content []byte) error {
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

func factoryDoGithub(method, url, token string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("call github err, url=%s, code=%d, res=%s", url, resp.StatusCode, bs)
	}

	return bs, nil
}

func getRepoPageURLInfo(repo, token string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/pages", repo)
	bs, err := factoryDoGithub(http.MethodGet, url, token)
	if err != nil {
		return "", err
	}

	var res = make(map[string]interface{})
	if err = json.Unmarshal(bs, &res); err != nil {
		return "", err
	}

	return res["html_url"].(string), nil
}

func getRepoPageBuildStatus(repo, token string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/pages/builds/latest", repo)
	bs, err := factoryDoGithub(http.MethodGet, url, token)
	if err != nil {
		return "", err
	}

	var res = make(map[string]interface{})
	if err = json.Unmarshal(bs, &res); err != nil {
		return "", err
	}

	return res["status"].(string), nil
}
