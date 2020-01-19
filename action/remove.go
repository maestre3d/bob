package action

import (
	"errors"

	"github.com/maestre3d/bob/usecase"
	"github.com/urfave/cli"
)

// RemoveApp Remove app
func RemoveApp(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return errors.New("Missing application name")
	}

	return usecase.RemoveApp(ctx.Args().First())
}

// RemoveService Remove service
func RemoveService(ctx *cli.Context) error {
	if ctx.NArg() < 2 {
		return errors.New("Missing application/service name")
	}

	return usecase.RemoveService(ctx.Args().First(), ctx.Args().Get(1))
}
