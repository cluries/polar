/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 * Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package api

import (
	"iusworks.com/p/fluence/web/route"

	verone "iusworks.com/p/polar/internal/web/action/api/v1"
)

var Routes = []*route.AppRoute{
	&verone.Route,
}
