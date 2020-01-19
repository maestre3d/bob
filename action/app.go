package action

import (
	"github.com/maestre3d/bob/usecase"
	"github.com/urfave/cli"
)

// GetAppInfo Get application(s) information
func GetAppInfo(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return usecase.GetAllAppInfo()
	}

	return usecase.GetAppInfo(ctx.Args().First())
}
