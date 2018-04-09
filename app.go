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

	"google.golang.org/appengine"
)

func main() {
	router := NewRouter()
	http.Handle("/", router)

	appengine.Main()
}
