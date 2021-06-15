/*
 * Project Polar
 *
 * We leave traces in this world ,
 * The times can not erase !
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 * Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package molecule

import (
	"net/http"

	"iusworks.com/p/fluence/web/action"
)

var Actions = []action.Action{
	{Name: "Auth", Method: http.MethodPost, Path: "auth", Action: authAction},
}
