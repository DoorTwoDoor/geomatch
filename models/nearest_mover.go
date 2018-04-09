/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package models provides data models for the application.
package models

// NearestMover represents a mover that is near a customer.
type NearestMover struct {
	Mover                  string  `json:"mover"`
	EstimatedTimeToArrival int     `json:"estimated_time_to_arrival"`
	Latitude               float64 `json:"latitude"`
	Longitude              float64 `json:"longitude"`
}
