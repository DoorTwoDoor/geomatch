/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec, parse requests
// write responses, validate, perform Cloud Datastore operations and perform
// Redis operations.
package utilities

import "log"

// Panic writes a panic log to the console in App Engine.
func Panic(value ...interface{}) {
	log.Panic(value)
}

// Print writes a print log to the console in App Engine.
func Print(value ...interface{}) {
	log.Print(value)
}
