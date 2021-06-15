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
	"github.com/gin-gonic/gin"
	"iusworks.com/p/fluence/web/action/api"
	"iusworks.com/p/fluence/web/tool"
)

//
func authAction(ctx *gin.Context) {
	var param = struct {
		Identifier string `form:"identifier" json:"identifier" xml:"identifier" binding:"required,min=4,max=32"`
		Credential string `form:"credential" json:"credential" xml:"credential"binding:"required,len=40,alphanum"`
		Captcha    string `form:"captcha" json:"captcha" xml:"captcha"binding:"required,len=8,alphanum"`
		Random     string `form:"random" json:"random" xml:"random" binding:"required,len=40,alphanum"`
	}{}

	if ex := tool.Bind(ctx, &param); ex != nil {
		api.Ex(ctx, ex)
		return
	}

	api.Ok(ctx, nil)
}
