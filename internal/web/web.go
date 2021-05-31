/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package web

import (
	"github.com/urfave/cli/v2"
	"iusworks.com/p/fluence/web/route"

	"iusworks.com/p/polar/internal/web/action/api"
)

func configureFluence(context *cli.Context) {
	 
}

func Loader(context *cli.Context) {
	configureFluence(context)

	for _, r := range api.Routes {
		route.Add(r)
	}
}
