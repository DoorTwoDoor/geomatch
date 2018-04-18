/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package handlers provides handler functions for the API endpoints.
package handlers

import (
	"net/http"

	"github.com/doortwodoor/geomatch/utilities"
)

// HandlePanic handles panics recovered from HTTP handlers.
func HandlePanic(
	responseWriter http.ResponseWriter,
	request *http.Request,
	value interface{},
) {
	shouldGzip := utilities.ShouldGzipResponse(request)

	utilities.WriteErrorResponse(
		responseWriter,
		http.StatusInternalServerError,
		shouldGzip,
	)
}
