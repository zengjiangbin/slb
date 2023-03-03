package main

import (
	"github.com/zengjiangbin/driver"
	"github.com/zengjiangbin/slb/core"
)

func main() {
	driver.Drive(core.NewService())
}
