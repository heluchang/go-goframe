package main

import (
	_ "test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
