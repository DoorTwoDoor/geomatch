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

	t.Run("active_mover", func(t *testing.T) {
		t.Parallel()

		onlineMover := OnlineMover{
			Move:      "0adiC7Dr5WBppb01Mjub",
			Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
			Latitude:  40.752556,
			Longitude: -73.977658,
			CreatedAt: createdAt,
		}

		assert.Equal(t, true, onlineMover.IsOnAMove())
	})

	t.Run("available_mover", func(t *testing.T) {
		t.Parallel()

		onlineMover := OnlineMover{
			Mover:     "KjfP77iiDSOKOoPEGnV0Jvmutcb2",
			Latitude:  30.452416,
			Longitude: -63.674854,
			CreatedAt: createdAt,
		}

		assert.Equal(t, false, onlineMover.IsOnAMove())
	})
}
