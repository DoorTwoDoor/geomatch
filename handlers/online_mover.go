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
	"strings"

	"github.com/doortwodoor/geomatch/models"
	"github.com/doortwodoor/geomatch/utilities"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

// PostOnlineMover creates or updates the location information for a mover.
func PostOnlineMover(
	validator utilities.Validator,
	redisClient utilities.RedisClient,
) httprouter.Handle {
	return func(
		responseWriter http.ResponseWriter,
		request *http.Request,
		_ httprouter.Params,
	) {
		// HTTP request header field names and values.
		const (
			acceptEncodingKey   = "Accept-Encoding"
			acceptEncodingValue = "gzip"
		)
		contentEncoding := request.Header.Get(acceptEncodingKey)
		shouldGzip := strings.Contains(contentEncoding, acceptEncodingValue)

		var error error
		onlineMover := models.OnlineMover{}
		error = utilities.Decode(request.Body, &onlineMover)
		if error != nil {
			utilities.WriteErrorResponse(
				responseWriter,
				http.StatusBadRequest,
				shouldGzip,
			)

			return
		}

		error = validator.ValidateStruct(onlineMover)
		if error != nil {
			utilities.WriteErrorResponse(
				responseWriter,
				http.StatusUnprocessableEntity,
				shouldGzip,
			)

			return
		}

		if onlineMover.IsOnAMove() { // Is the online mover on a move?
			// @TODO: Switch from datastore to Firestore.
			context := appengine.NewContext(request)
			kind := "OnlineMover"
			utilities.PutToDatastore(context, kind, &onlineMover)
		} else { // Is the online move available?
			key := "OnlineMovers"
			redisClient.GeoAdd(
				key,
				onlineMover.Mover,
				onlineMover.Latitude,
				onlineMover.Longitude,
			)
		}

		utilities.WriteOKResponse(responseWriter, onlineMover, shouldGzip)
	}
}
