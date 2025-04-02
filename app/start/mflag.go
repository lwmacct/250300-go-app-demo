package start

import (
	"github.com/lwmacct/250300-go-app-demo/app"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
	"github.com/spf13/cobra"
)

func Cmd() *mflag.Ts {
	mc := mflag.New(app.Flag).UsePackageName("")
	mc.AddCmd(func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	}, "run", "", "app", "mlog")
	return mc
}

func run(cmd *cobra.Command, args []string) {
	_ = map[string]any{"cmd": cmd, "args": args}
	mlog.Info(mlog.H{"msg": "app.Flag", "data": app.Flag})

	mlog.Close()
}
