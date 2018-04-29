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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsOnAMove(t *testing.T) {
	createdAt := time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC)
	onlineMoverOne := OnlineMover{
		Move:      "0adiC7Dr5WBppb01Mjub",
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: createdAt,
	}
	onlineMoverTwo := OnlineMover{
		Mover:     "KjfP77iiDSOKOoPEGnV0Jvmutcb2",
		Latitude:  30.452416,
		Longitude: -63.674854,
		CreatedAt: createdAt,
	}
	tests := []struct {
		testName       string
		testInput      OnlineMover
		expectedResult bool
	}{
		{"active_mover", onlineMoverOne, true},
		{"available_mover", onlineMoverTwo, false},
	}

	for _, test := range tests {
		test := test

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			expectedResult := test.expectedResult
			actualResult := test.testInput.IsOnAMove()

			assert.Equal(t, expectedResult, actualResult)
		})
	}
}
