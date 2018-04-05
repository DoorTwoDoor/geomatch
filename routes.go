/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"./handlers"

	"github.com/julienschmidt/httprouter"
)

// Route ...
type Route struct {
	Method  string
	Path    string
	Handler httprouter.Handle
}

// Routes ...
type Routes []Route

// RouterRoutes ...
var RouterRoutes = Routes{
	Route{
		Method:  "GET",
		Path:    "/nearest-movers",
		Handler: handlers.GetNearestMovers,
	},
	Route{
		Method:  "POST",
		Path:    "/online-movers",
		Handler: handlers.PostLocationUpdate,
	},
}
