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

package define

import "iusworks.com/p/fluence/tool/enum"

const platform = "Platform"

var (
	PlatformMobile = enum.NewBasicEnum(platform, 10, "Mobile")
	PlatformWechat = enum.NewBasicEnum(platform, 20, "Wechat")
	PlatformAlipay = enum.NewBasicEnum(platform, 30, "Alipay")
	PlatformApple  = enum.NewBasicEnum(platform, 40, "Apple")
)
