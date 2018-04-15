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

	"github.com/doortwodoor/geomatch/utilities"
	"google.golang.org/appengine"
)

func main() {
	// address := "35.184.145.206:6379"
	// password := "nY7FX1pqggQL"
	// if appengine.IsDevAppServer() { // Is running in the development app server?
	address := "localhost:6379"
	password := ""
	//}
	redisClient := utilities.NewRedisClient(address, password)
	router := NewRouter(redisClient)
	http.Handle("/", router)

	appengine.Main()
}
