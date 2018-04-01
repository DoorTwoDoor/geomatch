/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// package main

// import (
// 	"net/http"
// 	"strings"
// )

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	message := r.URL.Path
// 	message = strings.TrimPrefix(message, "/")
// 	message = "Hello " + message

// 	w.Write([]byte(message))
// }

// func main() {
// 	http.HandleFunc("/", sayHello)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"./models"
)

func main() {
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
	fmt.Println(string(b))

	timestamp, _ := time.Parse(time.RFC3339, "2018-03-06T04:31:45Z")
	location := models.LocationUpdate{
		Move:      "0adiC7Dr5WBppb01Mjub",
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  40.752556,
		Longitude: -73.977658,
		CreatedAt: timestamp,
	}
	c, _ := json.Marshal(location)
	fmt.Println(string(c))
}
