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
