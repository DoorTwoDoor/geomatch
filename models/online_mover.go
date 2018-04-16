/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package models provides data models for the application.
package models

import (
	"time"
)

// OnlineMover represents a mover that is either available or on a move.
type OnlineMover struct {
	Move      string    `json:"move,omitempty"`
	Mover     string    `json:"mover" validate:"required"`
	Latitude  float64   `json:"latitude" validate:"required,min=-90,max=90"`
	Longitude float64   `json:"longitude" validate:"required,min=-180,max=180"`
	CreatedAt time.Time `json:"created_at" validate:"required,lte"`
}

// IsOnAMove checks if an online mover is on a move.
func (onlineMover OnlineMover) IsOnAMove() bool {
	return onlineMover.Move != ""
}
