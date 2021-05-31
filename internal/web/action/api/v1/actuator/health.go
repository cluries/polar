/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 * Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package actuator

import (
	"github.com/gin-gonic/gin"
	"iusworks.com/p/fluence/web/action/api"

	"iusworks.com/p/polar/internal/web/action/api/v1/actuator/message"
)

func healthAction(ctx *gin.Context) {
	api.Ok(ctx, message.Healthy())
}
