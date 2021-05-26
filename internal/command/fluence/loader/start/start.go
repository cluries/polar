/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package start

import (
	"github.com/urfave/cli/v2"

	"iusworks.com/p/polar/internal/web"
)

func Loader(context *cli.Context) error {
	web.Loader(context)
	return nil
}
