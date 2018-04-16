/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec, write
// responses, perform Cloud Datastore operations and perform Redis operations.
package utilities

import "gopkg.in/go-playground/validator.v9"

// Validator represents a validator which validates values for structs and
// individual fields based on tags.
type Validator struct {
	validate *validator.Validate
}

// NewValidator returns a validator with sane defaults.
func NewValidator() Validator {
	validate := validator.New()

	return Validator{validate: validate}
}

// ValidateStruct validates a struct exposed fields, and automatically
// validates
func (validator Validator) ValidateStruct(value interface{}) error {
	return validator.validate.Struct(value)
}
