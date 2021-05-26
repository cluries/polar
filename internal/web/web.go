/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package web

import (
	"regexp"

	"github.com/urfave/cli/v2"
	"iusworks.com/p/fluence/exception"
	"iusworks.com/p/fluence/tool/stringx"
	"iusworks.com/p/fluence/web/middleware/limit"
	"iusworks.com/p/fluence/web/middleware/signature"
	"iusworks.com/p/fluence/web/route"

	"iusworks.com/p/polar/internal/web/action/api"
)

func configureFluence() {
	configureSignature()
	configureLimit()
}

func configureSignature() {
	sigconf := signature.GetConfig()
	sigconf.Lifeseconds = 300
	sigconf.PreliminaryValidor.Iden = func(iden string) bool {
		return regexp.MustCompile(`^\d{1,14}$`).MatchString(iden)
	}

	sigconf.TokenLoader = func(iden string) (string, *exception.AppException) {
		return stringx.SHA1(iden), nil
	}

	sigconf.NonceRedis = &struct {
		Host     string
		Port     int
		Database int
	}{Host: "127.0.0.1", Port: 6379, Database: 0}

	sigconf.Skipdoor = &struct {
		Key   string
		Value string
	}{Key: "Polarwego", Value: "Polarwego"}
}

func configureLimit() {
	limit.ConfigureRedis("127.0.0.1", 6379, 0)
}

func Loader(context *cli.Context) {
	configureFluence()

	for _, r := range api.Routes {
		route.Add(r)
	}
}
