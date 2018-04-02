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

	"../models"
	"../utilities"

	"github.com/julienschmidt/httprouter"
)

// PostLocationUpdate ...
func PostLocationUpdate(
	responseWriter http.ResponseWriter,
	request *http.Request,
	_ httprouter.Params,
) {
	// timestamp, _ := time.Parse(time.RFC3339, "2018-03-06T04:31:45Z")
	// location := models.LocationUpdate{
	// 	Move:      "0adiC7Dr5WBppb01Mjub",
	// 	Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
	// 	Latitude:  40.752556,
	// 	Longitude: -73.977658,
	// 	CreatedAt: timestamp,
	// }

	location := models.LocationUpdate{}
	utilities.Decode(request.Body, &location)

	if location.HasMoveID() {
		// Write data to Redis.
		fmt.Println("Mover is currently on an active move.")
	} else {
		// Write data to Google Cloud Datastore.
		fmt.Println("Mover is not currently on an active move.")
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	utilities.Encode(responseWriter, &location)
}
