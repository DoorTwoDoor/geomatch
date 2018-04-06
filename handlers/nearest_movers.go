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
)

// GetNearestMovers retrieves a set of movers that can travel to the
// specified location in the shortest amount of time.
func GetNearestMovers(
	responseWriter http.ResponseWriter,
	request *http.Request,
	_ httprouter.Params,
) {
	// Query Redis for movers within a certain radius of the location.
	// Use Google Time Distance Matrix API to calculate the estimated time to
	// arrival of each mover.
	mover := models.NearestMover{
		Mover: "5uls4pSbGeNvQFUYW8X74WraYcx2",
		EstimatedTimeToArrival: 122,
		Latitude:               40.752556,
		Longitude:              -73.977658,
	}
	// movers := models.NearestMovers{Movers: []models.NearestMover{}}
	// movers.Movers = []models.NearestMover{}
	var movers models.NearestMovers
	movers.Movers = append(movers.Movers, mover)

	contentEncoding := request.Header.Get("Accept-Encoding")
	shouldGzip := strings.Contains(contentEncoding, "gzip")

	utilities.WriteOKResponse(responseWriter, movers, shouldGzip)
}
