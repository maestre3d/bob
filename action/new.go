package action

import (
	"errors"
	"github.com/maestre3d/bob/usecase"
	"github.com/urfave/cli"
)

// NewApp Create a new app
func NewApp(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return errors.New("Missing new application name")
	}

	return usecase.CreateApp(ctx.Args().First(), ctx.Args().Get(1), ctx.String("d"))
}
