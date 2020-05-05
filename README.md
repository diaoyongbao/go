go
----------
# go语言学习
<!-- TOC depthFrom:2 orderedList:true -->

1. [go vscode配置](#go-vscode配置)
2. [go 基本](#go-基本)
3. [go mod 使用](#go-mod-使用)

<!-- /TOC -->
## go 在linux下的配置
1. 解压缩至/usr/local/目录下
2. 编辑系统环境文件/etc/profile
3. 添加如下环境变量
   export GOROOT=/usr/local/go
   export PATH=$PATH:$GOROOT/bin
4. 重新生效环境变量文件source /etc/profile
5. 使用go version 查看是否生效
## go vscode配置
1. go 的插件下载，vscode的插件中搜索go即可
2. go proxy代理使用go env -w GO111MODULE=on
   go env -w GOPROXY=https://goproxy.io,direct
3. tools安装，ctrl+shift+p 搜索go tools，全选安装
4. gocode代码提示插件，go get -u -v   github.com/stamblerre/gocode
5. 如果有失败的情况，可重启vscode再进行插件的安装
## go 基本
### go语言中的一些环境变量
GOROOT go的安装目录文件夹。如本文中将go安装在/esr/local/go文件下，此处的GOROOT=/usr/local/go
GOPROXY go的代理
GOPATH go的包及程序目录使用的文件下，此文件夹下存放go的包及插件
## go mod 使用
