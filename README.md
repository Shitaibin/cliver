cliver(CLI version demo)
=========

展示如何通过go build时动态的指定版本信息。

1. 设置环境变量。
```
VERSION=0.1
GIT_COMMIT=$(shell git rev-list -1 HEAD)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
```

2. go build时通过`ldflags`传递变量pkg的具体变量。

```
go build -ldflags "-X github.com/shitaibin/cliver/version.Version=$(VERSION)  -X github.com/shitaibin/cliver/version.GitCommit=$(GIT_COMMIT) -X github.com/shitaibin/cliver/version.BuildDate=$(BUILD_DATE)" .
```

注意：
i. 当前项目的完整名称已在go.mod中设置，所以`-X pkg.Var`中的pkg应当使用完整的包名称，否则无法设置成功。
ii. version包中提供了生成版本的函数`version.Get()`供`cliver version`命令调用。

3. 运行程序
```
$ ./cliver version
{Version: 0.1, GitCommit: 3977697acc47be3bc56f323847327ac7a13ad568, BuildDate: 2020-11-17T08:30:28Z, GoVersion: go1.15, Platform: darwin/amd64}
```
