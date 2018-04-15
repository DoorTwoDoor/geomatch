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
	Move      string    `json:"move,omitempty"`
	Mover     string    `json:"mover"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}

// IsOnAMove checks if an online mover is on a move.
func (onlineMover OnlineMover) IsOnAMove() bool {
	return onlineMover.Move != ""
}
