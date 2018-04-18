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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	const expectedResult = 10

	actualResult, _ := Atoi("10")

	assert.Equal(t, expectedResult, actualResult)
}

func TestParseFloat(t *testing.T) {
	const expectedResult = 43.474307

	actualResult, _ := ParseFloat("43.474307", 64)

	assert.Equal(t, expectedResult, actualResult)
}
