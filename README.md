#### presto go curd api  

#### go 版本
1.14


#### 部署说明 
```
GOPATH必须包含你下过的github包

当前依赖:
go get -v github.com/go-sql-driver/mysql
go get -v github.com/prestodb/presto-go-client/presto

# 项目目录放到gopath下
go env -w GO111MODULE=auto
go env -w GOPATH=$HOME/go/:~/go
go env -w GOPATH=$HOME/bin

# 中国区可选，VPN可以不执行，默认找不到仓库会去github回源
go env -w GOPROXY=https://goproxy.io,direct


LINUX编译命令
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

MACOS编译命令
go build
```

------------------------------------------------

#### go version
1.14


#### deployment
```
GOPATH

当前依赖:
go get -v github.com/go-sql-driver/mysql
go get -v github.com/prestodb/presto-go-client/presto

# GOPATH direct to your go src path, example: ~/go
go env -w GO111MODULE=auto
go env -w GOPATH=$HOME/go/:~/go
go env -w GOPATH=$HOME/bin



LINUX build bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

MACOS build bin
go build
```
