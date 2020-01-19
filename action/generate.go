package action

import (
	"errors"

	global "github.com/maestre3d/bob/common/config"
	"github.com/maestre3d/bob/usecase"
	"github.com/urfave/cli"
)

// GenerateService Create a new service
func GenerateService(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return errors.New(global.ServiceError)
	}

	return usecase.GenerateService(ctx.Args().First(), ctx.String("a"), ctx.String("d"))
}
