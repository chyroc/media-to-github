# media-to-github

上传小文件资源到 github，并利用 github pages 功能生成 url 。

## 安装
```
go get -u github.com/chyroc/media-to-github
```

## 使用

```text
NAME:
   media-to-github - 上传图片到github

USAGE:
   media-to-github [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -t value, --token value  github token
   -r value, --repo value   github repo
   -f value, --file value   file path or url
   -p value, --path value   where file store
   --help, -h               show help
   --version, -v            print the version
```

```text
// 从链接上传
media-to-github -t <github token> -r chyroc/media -f https://media.chyroc.cn/img/csrf.jpg -p data/test2.jpg

// 从本地文件上传
media-to-github -t <github token> -r chyroc/media -f ./README.md -p data/test3.md
```

`-p` 可以省略，将随机生成文件名（ `png` 后缀）
