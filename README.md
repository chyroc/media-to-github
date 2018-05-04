# media-to-github
上传img资源到github

## 安装
```
go get -u github.com/Chyroc/media-to-github
```

## 使用
```
// 从链接上传
media-to-github -t <github token> -r chyroc/media -f https://media.chyroc.cn/img/csrf.jpg -p img/test2.jpg

// 从本地文件上传
media-to-github -t <github token> -r chyroc/media -f ./README.md -p img/test3.md
```