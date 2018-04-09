/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import "net/http"

func main() {
	router := NewRouter()

	if err := http.ListenAndServe(":4040", router); err != nil {
		panic(err)
	}
}
