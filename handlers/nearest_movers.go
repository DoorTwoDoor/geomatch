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
		shouldGzip := utilities.ShouldGzipResponse(request)

		queryParameters, err := decodeQueryParameters(request)
		if err != nil {
			utilities.WriteErrorResponse(
				responseWriter,
				http.StatusBadRequest,
				shouldGzip,
			)

			return
		}

		err = validateQueryParameters(validator, queryParameters)
		if err != nil {
			utilities.WriteErrorResponse(
				responseWriter,
				http.StatusUnprocessableEntity,
				shouldGzip,
			)

			return
		}

		const (
			key             = "OnlineMovers"
			unit            = "m"
			withCoordinates = true
			sort            = "ASC"
		)
		latitude := queryParameters["latitude"].(float64)
		longitude := queryParameters["longitude"].(float64)
		radius := queryParameters["radius"].(float64)
		count := queryParameters["limit"].(int)

		// Query Redis for movers within a certain radius of the location.
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

		utilities.WriteOKResponse(responseWriter, movers, shouldGzip)
	}
}

// decodeQueryParameters decodes query parameters using decoding rules.
func decodeQueryParameters(
	request *http.Request,
) (map[string]interface{}, error) {
	decodingRules := map[string]string{
		"latitude":  "stof",
		"longitude": "stof",
		"radius":    "stof",
		"limit":     "stoi",
	}

	return utilities.DecodeQueryParameters(request, decodingRules)
}

// validateQueryParameters validates query parameters using tag style
// validation rules.
func validateQueryParameters(
	validator utilities.Validator,
	queryParameters map[string]interface{},
) error {
	validationRules := map[string]string{
		"latitude":  "required,min=-90,max=90",
		"longitude": "required,min=-180,max=180",
		"radius":    "required,min=0",
		"limit":     "required,min=1",
	}

	return validator.ValidateQueryParameters(queryParameters, validationRules)
}
