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

var generator = cli.Command{
	Name:        "generator",
	Description: "generator for config",
	Action:      generatorAction,
	Flags: []cli.Flag{
		&configFlag,
	},
}

func generatorAction(ctx *cli.Context) error {
	return nil
}
