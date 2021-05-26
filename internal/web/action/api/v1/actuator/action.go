/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package actuator

import (
	"net/http"

	"iusworks.com/p/fluence/web/action"
)

var Actions = []action.Action{
	{Name: "health", Method: http.MethodPost, Path: "health", Action: healthAction},
}
