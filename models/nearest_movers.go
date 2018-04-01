/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package models provides data models for the application.
package models

// NearestMovers represents a group of movers that are near a customer.
type NearestMovers struct {
	Movers []NearestMover `json:"nearest_movers"`
}
