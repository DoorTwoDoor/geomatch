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

import (
	"testing"
	"time"

	"github.com/doortwodoor/geomatch/models"
	"github.com/stretchr/testify/assert"
)

func TestNewValidator(t *testing.T) {
	validator := NewValidator()

	assert.NotNil(t, validator)
}

func TestValidateQueryParameters(t *testing.T) {
	queryParametersOne := map[string]interface{}{
		"latitude":  43.474307,
		"longitude": -80.537230,
		"radius":    1000,
		"limit":     10,
	}
	queryParametersTwo := map[string]interface{}{
		"latitude":  43.474307,
		"longitude": -180.537230,
		"radius":    1000,
		"limit":     10,
	}
	validationRules := map[string]string{
		"latitude":  "required,min=-90,max=90",
		"longitude": "required,min=-180,max=180",
		"radius":    "required,min=0",
		"limit":     "required,min=1",
	}
	tests := []struct {
		testName       string
		testsInputs    []interface{}
		expectedResult string
	}{
		{
			"valid_query_parameters",
			[]interface{}{queryParametersOne, validationRules},
			"nil",
		},
		{
			"invalid_query_parameters",
			[]interface{}{queryParametersTwo, validationRules},
			"not_nil",
		},
	}
	validator := NewValidator()

	for _, test := range tests {
		test := test

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			queryParameters := test.testsInputs[0].(map[string]interface{})
			rules := test.testsInputs[1].(map[string]string)
			result := validator.ValidateQueryParameters(queryParameters, rules)

			expectedResult := test.expectedResult

			if expectedResult == "nil" {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
			}
		})
	}
}

func TestValidateStruct(t *testing.T) {
	createdAt := time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC)
	validator := NewValidator()

	t.Run("active_mover", func(t *testing.T) {
		t.Parallel()

		onlineMover := models.OnlineMover{
			Move:      "0adiC7Dr5WBppb01Mjub",
			Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
			Latitude:  43.481082,
			Longitude: -80.530143,
			CreatedAt: createdAt,
		}
		result := validator.ValidateStruct(onlineMover)

		assert.Nil(t, result)
	})

	t.Run("available_mover", func(t *testing.T) {
		t.Parallel()

		onlineMover := models.OnlineMover{
			Mover:     "KjfP77iiDSOKOoPEGnV0Jvmutcb2",
			Latitude:  43.645621,
			Longitude: -79.391686,
			CreatedAt: createdAt,
		}
		result := validator.ValidateStruct(onlineMover)

		assert.Nil(t, result)
	})
}

func TestValidateVar(t *testing.T) {
	validator := NewValidator()

	t.Run("latitude", func(t *testing.T) {
		t.Parallel()
		const (
			value = 43.481082
			tag   = "required,min=-90,max=90"
		)
		result := validator.ValidateVar(value, tag)

		assert.Nil(t, result)
	})

	t.Run("longitude", func(t *testing.T) {
		t.Parallel()
		const (
			value = -80.530143
			tag   = "required,min=-180,max=180"
		)
		result := validator.ValidateVar(value, tag)

		assert.Nil(t, result)
	})

	t.Run("radius", func(t *testing.T) {
		t.Parallel()
		const (
			value = 1000
			tag   = "required,min=0"
		)
		result := validator.ValidateVar(value, tag)

		assert.Nil(t, result)
	})

	t.Run("limit", func(t *testing.T) {
		t.Parallel()
		const (
			value = 10
			tag   = "required,min=1"
		)
		result := validator.ValidateVar(value, tag)

		assert.Nil(t, result)
	})
}
