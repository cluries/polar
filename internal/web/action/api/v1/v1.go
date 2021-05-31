package v1

import (
	"github.com/gin-gonic/gin"
	"iusworks.com/p/fluence/web/action"
	"iusworks.com/p/fluence/web/middleware"
	"iusworks.com/p/fluence/web/middleware/limit"
	"iusworks.com/p/fluence/web/middleware/signature"
	"iusworks.com/p/fluence/web/route"

	"iusworks.com/p/polar/internal/web/action/api/v1/actuator"
)

var ipLimitor *limit.Limiter
var idenLimitor *limit.Limiter
var signator *signature.Signator

var Route = route.AppRoute{
	Name:       "apione",
	PathPrefix: "/api/v1",
	Actions: [][]action.Action{
		actuator.Actions,
	},
	Middlewares: []gin.HandlerFunc{
		middleware.Profile,
		ipLimitor.Handler,
		signator.Hander,
		idenLimitor.Handler,
	},
	Initializer: func() {
		ipLimitor = limit.NewLimitor(nil)
		idenLimitor = limit.NewLimitor(nil)
	},
}
