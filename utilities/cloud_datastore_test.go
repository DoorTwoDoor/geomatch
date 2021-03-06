/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec, parse requests
// write responses, validate, perform Cloud Datastore operations and perform
// Redis operations.
package utilities

import (
	"testing"
	"time"

	"github.com/doortwodoor/geomatch/models"
	"google.golang.org/appengine/aetest"
)

func TestPutToDatastore(t *testing.T) {
	context, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	const kind = "OnlineMover"
	onlineMover := models.OnlineMover{
		Move:      "0adiC7Dr5WBppb01Mjub",
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  40.752556,
		Longitude: -73.977658,
		CreatedAt: time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC),
	}
	_, err = PutToDatastore(context, kind, &onlineMover)
	if err != nil {
		t.Fatal(err)
	}
}
