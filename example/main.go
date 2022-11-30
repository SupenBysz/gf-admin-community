package main

import (
	"github.com/SupenBysz/gf-admin-community/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/SupenBysz/gf-admin-community"
)

func main() {
	cmd.Main.Run(gctx.New())
}
