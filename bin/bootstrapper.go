package bin

import (
	"fmt"
	"os"

	"github.com/maestre3d/bob/action"
	global "github.com/maestre3d/bob/common/config"
	"github.com/urfave/cli"
)

// Bootstrap Application wrapper
type Bootstrap struct{}

// InitCLI Initialize Command-Line Interface
func (b *Bootstrap) InitCLI() {
	// Create and configure CLI
	app := &cli.App{
		Name:         global.AppName,
		BashComplete: cli.DefaultAppComplete,
		Action: func(c *cli.Context) error {
			fmt.Printf("Welcome to %s, please select a command to get started. \n(Use -%s help- or -%s h- to get info)\n",
				global.AppName, global.ShortName, global.ShortName)
			return nil
		},
		Commands: []cli.Command{
			{
				// NEW
				Name:        "new",
				Aliases:     []string{"n"},
				Description: "Generate a new app",
				Usage:       "Generate new app",
				UsageText:   global.ShortName + " new {app_name} {path} \"{app_description}\"",
				ArgsUsage:   global.ShortName + " new foo",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "description, d",
						Value: "",
						Usage: "Description for app",
					},
				},
				Action: func(c *cli.Context) error {
					return action.NewApp(c)
				},
			},
			{
				// REMOVE
				Name:        "remove",
				Aliases:     []string{"rm"},
				Description: "Remove an existing module",
				Usage:       "Remove module",
				UsageText:   global.ShortName + " remove {command}",
				ArgsUsage:   global.ShortName + " remove service",
				Subcommands: cli.Commands{
					{
						Name:        "app",
						Aliases:     []string{"a"},
						Description: "Remove an existing app",
						Usage:       "Remove app",
						UsageText:   global.ShortName + " remove app {app_name}",
						ArgsUsage:   global.ShortName + " remove app foo",
						Action: func(c *cli.Context) error {
							return action.RemoveApp(c)
						},
					},
					{
						Name:        "service",
						Aliases:     []string{"s"},
						Description: "Remove an existing service",
						Usage:       "Remove service",
						UsageText:   global.ShortName + " remove service {app_name} {service_name}",
						ArgsUsage:   global.ShortName + " remove service foo bar",
						Action: func(c *cli.Context) error {
							return action.GenerateService(c)
						},
					},
				},
			},
			{
				// GENERATE
				Name:        "generate",
				Aliases:     []string{"g"},
				Description: "Create a new module",
				Usage:       "Generate a module",
				UsageText:   global.ShortName + " generate -command-",
				Subcommands: cli.Commands{
					{
						Name:        "service",
						Aliases:     []string{"s"},
						Description: "Create a new service",
						Usage:       "Generate a service",
						UsageText:   global.ShortName + " generate service {name} --app foo \"{service_description}\"",
						ArgsUsage:   global.ShortName + " generate service bar",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "app, a",
								Value: "",
								Usage: "Service's application",
							},
							&cli.StringFlag{
								Name:  "description, d",
								Value: "",
								Usage: "Description for service",
							},
						},
						Action: func(c *cli.Context) error {
							return action.GenerateService(c)
						},
					},
				},
			},
			{
				// APP
				Name:        "app",
				Aliases:     []string{"a"},
				Description: "Application operations",
				Usage:       global.ShortName + " app [command]",
				Subcommands: []cli.Command{
					{
						Name:        "info",
						Aliases:     []string{"i"},
						Description: "Show application(s) info",
						Usage:       "Show application(s)",
						UsageText:   global.ShortName + " app info {app_name}",
						ArgsUsage:   global.ShortName + " app info foo",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "app, a",
								Value: "",
								Usage: "Service's application",
							},
						},
						Action: func(c *cli.Context) error {
							return action.GetAppInfo(c)
						},
					},
				},
			},
			{
				// SERVICE
				Name:        "service",
				Aliases:     []string{"s"},
				Description: "Service operations",
				Usage:       global.ShortName + " service [command]",
				Subcommands: []cli.Command{
					{
						Name:        "info",
						Aliases:     []string{"i"},
						Description: "Show service(s) info",
						Usage:       "Show service(s)",
						UsageText:   global.ShortName + " service info {app_name} {service_name}",
						ArgsUsage:   global.ShortName + " service info foo bar",
						Action: func(c *cli.Context) error {
							return action.GetServiceInfo(c)
						},
					},
				},
			},
		},
		Description:          global.AppName + " is a Microservice Builder that simplifies the process of making a service from the scratch.",
		Author:               "Alonso R",
		Email:                "aruiz@damascus-engineering.com",
		EnableBashCompletion: true,
		Version:              "1.0",
		Copyright:            "This tool was developed by Alonso R -maestre3d- and is under the MIT license. (2020)",
	}

	// Run CLI
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf(global.GenericError, err.Error())
	}
}
