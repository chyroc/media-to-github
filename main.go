package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/Chyroc/media-to-github/internal"
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
	app.Usage = "上传图片到github"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "t, token",
			Usage:       "github token",
			Destination: &token,
		},
		cli.StringFlag{
			Name:        "r, repo",
			Usage:       "github repo",
			Destination: &repo,
		},
		cli.StringFlag{
			Name:        "f, file",
			Usage:       "file path or url",
			Destination: &file,
		},
		cli.StringFlag{
			Name:        "p, path",
			Usage:       "where file store",
			Destination: &path,
		},
	}

	app.Action = func(c *cli.Context) error {
		if token == "" || repo == "" || file == "" || path == "" {
			return cli.ShowAppHelp(c)
		}

		var content []byte

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
			content = []byte(file)
		}

		return internal.Upload(repo, token, path, content)
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
