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
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestNewRedisClient(t *testing.T) {
	const (
		address  = "localhost:6379"
		password = ""
	)
	redisClient := NewRedisClient(address, password)

	assert.NotNil(t, redisClient)
}

func TestGeoAdd(t *testing.T) {
	const expectedResult = int64(1)

	const (
		address  = "localhost:6379"
		password = ""
	)
	redisClient := NewRedisClient(address, password)

	const key = "OnlineMovers"
	onlineMover := models.OnlineMover{
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC),
	}
	redisClient.ZRem(key, onlineMover.Mover)

	actualResult, _ := redisClient.GeoAdd(
		key,
		onlineMover.Mover,
		onlineMover.Latitude,
		onlineMover.Longitude,
	)

	assert.Equal(t, expectedResult, actualResult)
}

func TestGeoRadius(t *testing.T) {
	expectedResult := []redis.GeoLocation{
		redis.GeoLocation{
			Name:      "5uls4pSbGeNvQFUYW8X74WraYcx2",
			Latitude:  43.48108202751774,
			Longitude: -80.53014189004898,
			Dist:      0,
			GeoHash:   0,
		},
	}

	const (
		address  = "localhost:6379"
		password = ""
	)
	redisClient := NewRedisClient(address, password)

	createdAt := time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC)
	onlineMoverOne := models.OnlineMover{
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: createdAt,
	}
	onlineMoverTwo := models.OnlineMover{
		Mover:     "KjfP77iiDSOKOoPEGnV0Jvmutcb2",
		Latitude:  43.645621,
		Longitude: -79.391686,
		CreatedAt: createdAt,
	}

	const key = "OnlineMovers"
	redisClient.GeoAdd(
		key,
		onlineMoverOne.Mover,
		onlineMoverOne.Latitude,
		onlineMoverOne.Longitude,
	)
	redisClient.GeoAdd(
		key,
		onlineMoverTwo.Mover,
		onlineMoverTwo.Latitude,
		onlineMoverTwo.Longitude,
	)

	const (
		latitude        = 43.474307
		longitude       = -80.537230
		radius          = 1000
		unit            = "m"
		withCoordinates = true
		count           = 10
		sort            = "ASC"
	)
	actualResult, _ := redisClient.GeoRadius(
		key,
		latitude,
		longitude,
		radius,
		unit,
		withCoordinates,
		count,
		sort,
	)

	assert.Equal(t, expectedResult, actualResult)
}

func TestZRem(t *testing.T) {
	const expectedResult = int64(1)

	const (
		address  = "localhost:6379"
		password = ""
	)
	redisClient := NewRedisClient(address, password)

	const key = "OnlineMovers"
	onlineMover := models.OnlineMover{
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC),
	}
	redisClient.GeoAdd(
		key,
		onlineMover.Mover,
		onlineMover.Latitude,
		onlineMover.Longitude,
	)

	actualResult, _ := redisClient.ZRem(key, onlineMover.Mover)

	assert.Equal(t, expectedResult, actualResult)
}
