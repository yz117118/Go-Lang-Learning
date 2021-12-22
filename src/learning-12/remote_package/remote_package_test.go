package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRPC(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
	t.Log(m.Get(cm.StrKey("key1")))
}

//remote package
//go get or go get -u
//go source push to github
//注意代码在GitHub上的组织形式，以适应go get
//直接以代码路径开始，不要有src

//sync map 在读很多写很少时，性能很高
//go env -w GO111MODULE=on
//go env -w GOPROXY=https://goproxy.io,direct
//#设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
//go env -w GOPRIVATE=*.gitlab.com
// go get -u github.com/easierway/concurrent_map"
