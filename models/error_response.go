/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package models provides data models for the application.
package models

// ErrorResponse represents an error returned by the application.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
