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

	"./handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nearest-movers", handlers.GetNearestMovers)
	router.POST("/online-movers", handlers.PostLocationUpdate)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
