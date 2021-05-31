/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 * Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package configure

//
//func Middlewares() []func(context *cli.Context) {
//	const configures = 2

	//var functions = make([]func(context *cli.Context), configures)
	//
	//functions[0] = func(context *cli.Context) {
	//	sigconf := signature.GetConfig()
	//	sigconf.Lifeseconds = 300
	//
	//	sigconf.NonceRedis = &struct {
	//		Host     string
	//		Port     int
	//		Database int
	//	}{Host: "127.0.0.1", Port: 6379, Database: 0}
	//
	//	sigconf.Skipdoor = &struct {
	//		Key   string
	//		Value string
	//	}{Key: "Polarwego", Value: "Polarwego"}
	//
	//	sigconf.PreliminaryValidor.Iden = func(iden string) bool {
	//		return regexp.MustCompile(`^\d{1,14}$`).MatchString(iden)
	//	}
	//
	//	sigconf.TokenLoader = func(iden string) (string, *exception.AppException) {
	//		return stringx.SHA1(iden), nil
	//	}
	//
	//	sigconf.AllRightor = func(sigctx *signature.SigContext, ginctx *gin.Context) *exception.AppException {
	//		return nil
	//	}
	//}

//
//	return functions
//}
