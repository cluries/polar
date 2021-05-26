package v1

import (
	"github.com/gin-gonic/gin"
	"iusworks.com/p/fluence/web/action"
	"iusworks.com/p/fluence/web/middleware"
	"iusworks.com/p/fluence/web/middleware/limit"
	"iusworks.com/p/fluence/web/route"

	"iusworks.com/p/polar/internal/web/action/api/v1/actuator"
)

var Route = route.AppRoute{
	Name:       "apione",
	PathPrefix: "/api/v1",
	Actions: [][]action.Action{
		actuator.Actions,
	},
	Middlewares: []gin.HandlerFunc{
		middleware.Profile,
		middleware.Limit(&limit.Config{
			Cycle:       60,
			Limit:       60,
			IdenFetcher: limit.IPIdenFetcher,
		}),
		middleware.Signature,
		middleware.Limit(&limit.Config{
			Cycle: 60,
			Limit: 60,
			IdenFetcher: func(ctx *gin.Context) (string, error) {
				return "1", nil
			},
		}),
	},
}
