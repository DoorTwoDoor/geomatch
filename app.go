/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./models"
	"github.com/julienschmidt/httprouter"
)

func getNearestMovers(w http.ResponseWriter, r *http.Request) {
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
	b, _ := json.Marshal(movers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getLocationUpdate(responseWriter http.ResponseWriter, request *http.Request) {
	// timestamp, _ := time.Parse(time.RFC3339, "2018-03-06T04:31:45Z")
	// location := models.LocationUpdate{
	// 	Move:      "0adiC7Dr5WBppb01Mjub",
	// 	Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
	// 	Latitude:  40.752556,
	// 	Longitude: -73.977658,
	// 	CreatedAt: timestamp,
	// }
	// c, _ := json.Marshal(location)

	location := models.LocationUpdate{}
	json.NewDecoder(request.Body).Decode(&location)
	c, _ := json.Marshal(location)

	if location.HasMoveID() {
		// Write data to Redis.
		fmt.Println("Mover is currently on an active move.")
	} else {
		// Write data to Google Cloud Datastore.
		fmt.Println("Mover is not currently on an active move.")
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(c)
}

func main() {
	router := httprouter.New()
	fmt.Println(router)
	http.HandleFunc("/nearest-movers", getNearestMovers)
	http.HandleFunc("/available-movers", getLocationUpdate)
	// http.HandleFunc("/unavailable-movers", removeMover)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
