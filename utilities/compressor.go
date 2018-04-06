/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec and write
// responses.
package utilities

import (
	"compress/gzip"
	"net/http"
)

// NewGzipWriter returns a new writer.
func NewGzipWriter(responseWriter http.ResponseWriter) *gzip.Writer {
	return gzip.NewWriter(responseWriter)
}
