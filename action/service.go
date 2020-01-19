package action

import (
	"errors"

	"github.com/maestre3d/bob/usecase"
	"github.com/urfave/cli"
)

// GetServiceInfo Get service(s) information
func GetServiceInfo(ctx *cli.Context) error {
	if ctx.NArg() < 2 {
		return usecase.GetAllServiceInfo(ctx.Args().First())
	} else if ctx.NArg() < 1 {
		return errors.New("Missing application name")
	}

	return usecase.GetServiceInfo(ctx.Args().First(), ctx.Args().Get(1))
}
