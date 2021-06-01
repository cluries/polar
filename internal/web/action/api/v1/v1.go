package v1

import (
	"github.com/gin-gonic/gin"
	"iusworks.com/p/fluence/config"
	"iusworks.com/p/fluence/exception"
	"iusworks.com/p/fluence/web/action"
	"iusworks.com/p/fluence/web/middleware"
	"iusworks.com/p/fluence/web/middleware/limit"
	"iusworks.com/p/fluence/web/middleware/signature"
	"iusworks.com/p/fluence/web/route"

	"iusworks.com/p/polar/internal/web/action/api/v1/actuator"
)

func Approute() *route.AppRoute {
	ipLimiter := limit.NewLimiter(nil)
	idenLimiter := limit.NewLimiter(nil)
	signator := signature.NewSignator(nil)

	r := route.AppRoute{
		Name:       "apione",
		PathPrefix: "/api/v1",
		Actions: [][]action.Action{
			actuator.Actions,
		},
		Middlewares: []gin.HandlerFunc{
			middleware.Profile,
			ipLimiter.Handler,
			signator.Hander,
			idenLimiter.Handler,
		},
		Initializer: func() {
			ipLimiter.Update(ipLimiterConfig())
			idenLimiter.Update(idenLimiterConfig())
			signator.Update(signatureConfig())
		},
	}

	return &r
}

func ipLimiterConfig() *limit.Config {
	p := config.SharedProperty().Datasource.Redis
	c := limit.Config{
		Name:  "IPLimiter",
		Cycle: 60,
		Limit: 60,
		Redis: struct {
			Host     string
			Port     int
			Database int
		}{Host: p.Host, Port: p.Port, Database: p.Database},
		IdenFetcher: limit.IPIdenFetcher,
	}
	return &c
}

func idenLimiterConfig() *limit.Config {
	p := config.SharedProperty().Datasource.Redis
	c := limit.Config{
		Name:  "IdenLimiter",
		Cycle: 60,
		Limit: 60,
		Redis: struct {
			Host     string
			Port     int
			Database int
		}{Host: p.Host, Port: p.Port, Database: p.Database},
		IdenFetcher: func(ctx *gin.Context) (string, error) {
			return "123", nil
		},
	}
	return &c
}

func signatureConfig() *signature.Config {
	p := config.SharedProperty().Datasource.Redis
	c := signature.Config{
		Name:           "APISignature",
		Lifeseconds:    300,
		NonceRedisPool: 1024,
	}

	c.NonceRedis = &struct {
		Host     string
		Port     int
		Database int
	}{Host: p.Host, Port: p.Port, Database: p.Database}

	c.Skipdoor = &struct {
		Key   string
		Value string
	}{Key: "Polarskip", Value: "Polarskip"}

	c.PreliminaryValidor.Iden = func(iden string) bool {
		return len(iden) >= 1
	}

	c.TokenLoader = func(iden string) (string, *exception.AppException) {
		return "6120180395d88561f4769b5b6ca01938e7f26211", nil
	}

	return &c
}
