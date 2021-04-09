package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"github.com/urfave/cli"
)

var version = "v0.1.0"

func main() {
	var token string
	var repo string
	var file string
	var path string

	app := cli.NewApp()
	app.Name = "media-to-github"
	app.Version = version
	app.HelpName = "media-to-github"
	app.Usage = "上传图片到 GitHub"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "t, token",
			Usage:       "github token(default: read from env GITHUB_TOKEN)",
			Destination: &token,
		},
		cli.StringFlag{
			Name:        "r, repo",
			Usage:       "github repo",
			Destination: &repo,
		},
		cli.StringFlag{
			Name:        "f, file",
			Usage:       "file path or url(default: data/<uuid>.png)",
			Destination: &file,
		},
		cli.StringFlag{
			Name:        "p, path",
			Usage:       "where file store(default: png file from parse)",
			Destination: &path,
		},
	}

	app.Action = func(c *cli.Context) error {
		if token == "" {
			token = os.Getenv("GITHUB_TOKEN")
		}
		if token == "" || repo == "" {
			return cli.ShowAppHelp(c)
		}
		if path == "" {
			path = fmt.Sprintf("data/%s.png", uuid.NewV4().String())
		}
		if file == "" {
			var err error
			file, err = readImageByParse()
			if err != nil {
				return err
			}
		}

		var content []byte
		var err error

		if strings.HasPrefix(file, "http") {
			resp, err := http.Get(file)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			content, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
		} else {
			content, err = ioutil.ReadFile(file)
			if err != nil {
				return err
			}
		}

		if err = upload(repo, token, path, content); err != nil {
			return err
		}

		pageURL, err := getRepoPageURLInfo(repo, token)
		if err != nil {
			return nil
		}

		fmt.Printf("url: %s%s\n", pageURL, path)

		for {
			status, err := getRepoPageBuildStatus(repo, token)
			if err != nil {
				return err
			}
			if status == "built" {
				fmt.Printf("built done.\n")
				return nil
			}
			fmt.Printf("build %s, wait 2s.\n", status)
			time.Sleep(time.Second * 2)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
