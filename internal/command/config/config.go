/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package config

import "github.com/urfave/cli/v2"

var Config = cli.Command{
	Name:        "config",
	Description: "config",
	Subcommands: []*cli.Command{
		&validator,
		&generator,
	},
}

var configFlag = cli.StringFlag{
	Name:    "c",
	Usage:   "config file",
	Value:   "./configs/application.yml",
	EnvVars: []string{"POLAR_CONFIG"},
}
