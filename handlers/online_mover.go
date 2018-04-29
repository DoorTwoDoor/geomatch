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
		shouldGzip := utilities.ShouldGzipResponse(request)

		onlineMover := models.OnlineMover{}
		err := utilities.Decode(request.Body, &onlineMover)
		if err != nil {
			utilities.Print(err)

			utilities.WriteErrorResponse(
				responseWriter,
				http.StatusBadRequest,
				shouldGzip,
			)

			return
		}

		err = validator.ValidateStruct(onlineMover)
		if err != nil {
			utilities.Print(err)

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
		} else { // Is the online mover available?
			key := "OnlineMovers"
			_, err := redisClient.GeoAdd(
				key,
				onlineMover.Mover,
				onlineMover.Latitude,
				onlineMover.Longitude,
			)
			if err != nil {
				utilities.Panic(err)
			}
		}

		utilities.WriteOKResponse(responseWriter, onlineMover, shouldGzip)
	}
}
