package main

import (
	_ "gf-admin/packed"

	cmd "gf-admin/app/system/admin"
	_ "gf-admin/boot"

	"github.com/gogf/gf/v2/os/gctx"
)

//初始化admin用户

func main() {

	ctx := gctx.New()
	cmd.DwRun(ctx)
	cmd.Run(ctx)
}
