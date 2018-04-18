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

import "strconv"

// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
func Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseFloat converts the string s to a floating-point number with the
// precision specified by bit size.
func ParseFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}
