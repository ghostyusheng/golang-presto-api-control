#### presto go curd api  

#### go 版本
1.14


#### 部署说明 
```
1.记得修改conn.go presto_dsn 和 mysql 连接配置
2.mysql 登录:
  CREATE DATABASE `go_interface`;
  CREATE TABLE `access` (
  `ak` varchar(255) NOT NULL,
  `tables` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`ak`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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

#### 增删查改/CURD API DEMO
```
curl -X POST \
  http://127.0.0.1:8088/create \
  -H 'Content-Type: application/json' \
  -d '{
	"dest": "hive",
	"table": "testDemo",
	"primary": "country",
	"option": "create",
	"fields": ["user_id", "view_time", "page_url", "ds", "country"],
	"ak": "xd_search"
}'

curl -X POST \
  http://127.0.0.1:8088/query \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"dest": "hive",
	"option": "query",
	"sql": "select vn_code,md from (select * from testDemo limit 100)t where t.vn_code ='\''624'\''",
	"ak": "xd_search"
}'

curl -X POST \
  http://127.0.0.1:8088/truncate \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"dest": "hive",
	"table": "testDemo",
	"option": "truncate",
	"ak": "xd_search"
}'

{
	"dest": "hive",
	"table": "testDemo",
	"option": "insert",
	"fields": ["user_id", "view_time", "page_url", "ds", "country"],
	"data": [["77777","2","3","4","us"],["77777","2b","3c","4d","china"]],
	"ak": "xd_search"
}

```
------------------------------------------------



#### go version
1.14


#### deployment
`firstly, edit conn.go presto_dsn and mysql connect config`
```
# LOGIN MYSQL SHELL AND EXECUTE:
  CREATE DATABASE `go_interface`;
  
  CREATE TABLE `access` (
  `ak` varchar(255) NOT NULL,
  `tables` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`ak`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

go get -v github.com/go-sql-driver/mysql
go get -v github.com/prestodb/presto-go-client/presto

# GOPATH direct to your go src path, example: ~/go
go env -w GO111MODULE=auto
go env -w GOPATH=$HOME/go/:~/go
go env -w GOPATH=$HOME/bin


# LINUX build bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# MACOS build bin
go build
```
