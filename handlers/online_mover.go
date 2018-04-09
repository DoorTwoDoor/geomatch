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
	"fmt"
	"net/http"
	"strings"

	"github.com/doortwodoor/geomatch/models"
	"github.com/doortwodoor/geomatch/utilities"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

// PostOnlineMover creates or updates the location information for a mover.
func PostOnlineMover(
	responseWriter http.ResponseWriter,
	request *http.Request,
	_ httprouter.Params,
) {
	// HTTP response header field names and values.
	const (
		acceptEncodingKey   = "Accept-Encoding"
		acceptEncodingValue = "gzip"
	)

	onlineMover := models.OnlineMover{}
	utilities.Decode(request.Body, &onlineMover)

	if onlineMover.IsOnAMove() { // Is the online mover on a move?
		// Write data to Google Cloud Datastore.
		fmt.Println("Mover is currently on an active move.")
		context := appengine.NewContext(request)
		utilities.PutToDatastore(context, "OnlineMover", &onlineMover)
	} else { // Is the online mover available?
		// Write data to Redis.
		fmt.Println("Mover is not currently on an active move.")
	}

	contentEncoding := request.Header.Get(acceptEncodingKey)
	shouldGzip := strings.Contains(contentEncoding, acceptEncodingValue)

	utilities.WriteOKResponse(responseWriter, onlineMover, shouldGzip)
}
