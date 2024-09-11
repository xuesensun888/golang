package internal

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type DebInfo struct {
	Package string `yaml:"package"`
	Status  string `yaml:"status"`
	Version string `yaml:"version"`
	Arch    string `yaml:"arch"`
}

// 获取deb包的状态信息
func DebQuery(name string) (debInfo *DebInfo) {
	debInfo = new(DebInfo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second)*10)
	defer cancel()
	cmd := fmt.Sprintf("dpkg-query -f='package: ${Package}\nstatus: ${db:Status-Status}\nversion: ${Version}\narch: ${Architecture}\n' -W %s", name)
	stdout, _, exitcode := CommandContext(ctx, cmd)
	if exitcode != 0 {
		debInfo.Version = "0.0.0"
		return
	}
	_ = yaml.Unmarshal([]byte(stdout), debInfo)
	return
}
