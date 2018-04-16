/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"net/http"

	"github.com/doortwodoor/geomatch/handlers"
	"github.com/doortwodoor/geomatch/utilities"
	"github.com/julienschmidt/httprouter"
)

// NewRouter returns a new initialized router.
func NewRouter(
	validator utilities.Validator,
	redisClient utilities.RedisClient,
) *httprouter.Router {
	router := httprouter.New()

	for _, route := range RouterRoutes {
		router.Handle(
			route.Method,
			route.Path,
			route.Handler(validator, redisClient),
		)
	}

	router.NotFound = http.HandlerFunc(handlers.HandleNotFound)
	router.MethodNotAllowed = http.HandlerFunc(handlers.HandleMethodNotAllowed)
	router.PanicHandler = handlers.HandlePanic

	return router
}
