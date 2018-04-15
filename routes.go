/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"github.com/doortwodoor/geomatch/handlers"
	"github.com/doortwodoor/geomatch/utilities"
	"github.com/julienschmidt/httprouter"
)

// Handler represents handler functions for the API endpoints.
type Handler func(redisClient utilities.RedisClient) httprouter.Handle

// Route represents an API endpoint that the application can handle.
type Route struct {
	Method  string
	Path    string
	Handler Handler
}

// Routes represents a list of API endpoints that the application can handle.
type Routes []Route

// HTTP methods and API endpoint paths.
const (
	getMethod         = "GET"
	postMethod        = "POST"
	nearestMoversPath = "/nearest-movers"
	onlineMoversPath  = "/online-movers"
)

// RouterRoutes stores the list of API endpoints that the application can
// handle.
var RouterRoutes = Routes{
	Route{
		Method:  getMethod,
		Path:    nearestMoversPath,
		Handler: handlers.GetNearestMovers,
	},
	Route{
		Method:  postMethod,
		Path:    onlineMoversPath,
		Handler: handlers.PostOnlineMover,
	},
}
