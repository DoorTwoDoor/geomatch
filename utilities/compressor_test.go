/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec, write
// responses, perform Cloud Datastore operations and perform Redis operations.
package utilities

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGzipWriter(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	gzipWriter := NewGzipWriter(responseWriter)

	assert.NotNil(t, gzipWriter)
}
