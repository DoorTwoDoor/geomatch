/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package models provides data models for the application.
package models

import "time"

// LocationUpdate represents a location update from a mover.
type LocationUpdate struct {
	Move      string    `json:"move"`
	Mover     string    `json:"mover"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}
