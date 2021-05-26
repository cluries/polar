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

import (
	"github.com/urfave/cli/v2"
)



var validator = cli.Command{
	Name:        "validator",
	Description: "validator for config",
	Action:      validatorAction,
	Flags: []cli.Flag{
		&configFlag,
	},
}

func validatorAction(ctx *cli.Context) error {
	return nil
}
