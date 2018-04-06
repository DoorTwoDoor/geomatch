/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import "github.com/julienschmidt/httprouter"

// NewRouter returns a new initialized router.
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range RouterRoutes {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	return router
}
