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

// OnlineMover represents a mover that is either available or on a move.
type OnlineMover struct {
	Move      string    `datastore:"move" json:"move,omitempty"`
	Mover     string    `datastore:"mover" json:"mover"`
	Latitude  float64   `datastore:"latitude" json:"latitude"`
	Longitude float64   `datastore:"longitude" json:"longitude"`
	CreatedAt time.Time `datastore:"createdAt" json:"created_at"`
}

// IsOnAMove checks if an online mover is on a move.
func (onlineMover OnlineMover) IsOnAMove() bool {
	return onlineMover.Move != ""
}
