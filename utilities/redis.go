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

import "github.com/go-redis/redis"

// RedisClient represents a Redis client with a pool of zero or more
// underlying connections.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient returns a client to the Redis server sepcified by options.
func NewRedisClient(address string, password string) RedisClient {
	options := redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	}
	client := redis.NewClient(&options)

	return RedisClient{client: client}
}

// GeoAdd adds the specified geospatial item (name, latitude, longitude) to the
// specified key.
func (redisClient RedisClient) GeoAdd(
	key string,
	name string,
	latitude float64,
	longitude float64,
) (int64, error) {
	geoLocation := redis.GeoLocation{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}

	return redisClient.client.GeoAdd(key, &geoLocation).Result()
}

// GeoRadius returns the members of a sorted set populated with geospatial
// information using GeoAdd, which are within the borders of the area
// specified with the center location and the maximum distance from the
// center (the radius).
func (redisClient RedisClient) GeoRadius(
	key string,
	latitude float64,
	longitude float64,
	radius float64,
	unit string,
	withCoordinates bool,
	count int,
	sort string,
) ([]redis.GeoLocation, error) {
	geoRadiusQuery := redis.GeoRadiusQuery{
		Radius:    radius,
		Unit:      unit,
		WithCoord: withCoordinates,
		Count:     count,
		Sort:      sort,
	}

	return redisClient.client.GeoRadius(key, longitude, latitude, &geoRadiusQuery).Result()
}

// ZRem removes the specified member from the sorted set stored at key.
// @NOTE: Non existing members are ignored.
func (redisClient RedisClient) ZRem(
	key string,
	member interface{},
) (int64, error) {
	return redisClient.client.ZRem(key, member).Result()
}
