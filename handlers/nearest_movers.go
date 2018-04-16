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
	"strconv"
	"strings"

	"github.com/doortwodoor/geomatch/models"
	"github.com/doortwodoor/geomatch/utilities"
	"github.com/julienschmidt/httprouter"
	// "googlemaps.github.io/maps"
)

// GetNearestMovers retrieves a set of movers that can travel to the
// specified location in the shortest amount of time.
func GetNearestMovers(
	validator utilities.Validator,
	redisClient utilities.RedisClient,
) httprouter.Handle {
	return func(
		responseWriter http.ResponseWriter,
		request *http.Request,
		_ httprouter.Params,
	) {
		// Query Redis for movers within a certain radius of the location.
		const (
			key             = "OnlineMovers"
			unit            = "m"
			withCoordinates = true
			sort            = "ASC"
		)
		queryStrings := request.URL.Query()
		latitude, _ := strconv.ParseFloat(queryStrings.Get("latitude"), 64)
		longitude, _ := strconv.ParseFloat(queryStrings.Get("longitude"), 64)
		radius, _ := strconv.ParseFloat(queryStrings.Get("radius"), 64)
		count, _ := strconv.Atoi(queryStrings.Get("limit"))

		// @TODO: Need to validate the query strings here after parsing.

		nearestMovers, _ := redisClient.GeoRadius(
			key,
			latitude,
			longitude,
			radius,
			unit,
			withCoordinates,
			count,
			sort,
		)

		// Use Google Time Distance Matrix API to calculate the estimated time to
		// arrival of each mover.

		movers := models.NearestMovers{Movers: make([]models.NearestMover, 0)}
		for _, mover := range nearestMovers {
			mover := models.NearestMover{
				Mover: mover.Name,
				EstimatedTimeToArrival: 122,
				Latitude:               mover.Latitude,
				Longitude:              mover.Longitude,
			}
			movers.Movers = append(movers.Movers, mover)
		}

		// HTTP request header field names and values.
		const (
			acceptEncodingKey   = "Accept-Encoding"
			acceptEncodingValue = "gzip"
		)
		contentEncoding := request.Header.Get(acceptEncodingKey)
		shouldGzip := strings.Contains(contentEncoding, acceptEncodingValue)

		utilities.WriteOKResponse(responseWriter, movers, shouldGzip)
	}
}
