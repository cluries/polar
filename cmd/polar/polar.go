/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package main

import (
	"github.com/urfave/cli/v2"
	"iusworks.com/p/fluence/app"

	"iusworks.com/p/polar/internal/command/config"
	"iusworks.com/p/polar/internal/command/fluence/loader/start"
)

const name = "Godzilla"

func main() {
	commands := []*cli.Command{
		&config.Config,
	}

	flu := app.MakeFluence(name, commands)
	flu.FluenceCommandLoader.Start = start.Loader

	flu.Run()
}
