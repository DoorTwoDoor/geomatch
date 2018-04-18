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

// ValidateQueryParameters validates query parameters using tag style
// validation rules.
func (validator Validator) ValidateQueryParameters(
	queryParameters map[string]interface{},
	rules map[string]string,
) error {
	var err error

	for key, value := range rules {
		queryParameter := queryParameters[key]

		err = validator.ValidateVar(queryParameter, value)

		if err != nil {
			break
		}
	}

	return err
}

// ValidateStruct validates a struct exposed fields, and automatically
// validates nested structs, unless otherwise specified.
func (validator Validator) ValidateStruct(value interface{}) error {
	return validator.validate.Struct(value)
}

// ValidateVar validates a single variable using tag style validation.
func (validator Validator) ValidateVar(value interface{}, tag string) error {
	return validator.validate.Var(value, tag)
}
