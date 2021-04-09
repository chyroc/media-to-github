# media-to-github

上传小文件资源到 github，并利用 github pages 功能生成 url 。

## Install

### By Go Get

```shell
go get github.com/chyroc/chyroc/media-to-github
```

### By Brew

```shell
brew tap chyroc/tap
brew install chyroc/tap/chyroc/media-to-github
```

## Usage

```text
NAME:
   media-to-github - 上传图片到 GitHub

USAGE:
   media-to-github [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -t value, --token value  github token(default: read from env GITHUB_TOKEN)
   -r value, --repo value   github repo
   -f value, --file value   file path or url(default: data/<uuid>.png)
   -p value, --path value   where file store(default: png file from parse)
   --help, -h               show help
   --version, -v            print the version
```

```text
// 从链接上传
media-to-github -t <github token> -r owner/repo -f https://host/csrf.jpg

// 从本地文件上传
media-to-github -t <github token> -r owner/repo -f ./README.md

// 先复制图片，然后从剪贴板粘贴
media-to-github -t <github token> -r owner/repo
```

- `-t` 可以省略，默认读取环境变量：GITHUB_TOKEN
- `-p` 可以省略，默认生成随机文件名（ `png` 后缀）
- `-f` 可以省略，默认从剪切板读图片
