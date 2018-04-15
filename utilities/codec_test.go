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

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/doortwodoor/geomatch/models"
	"github.com/stretchr/testify/assert"
)

// compactJSON returns a JSON-encoded string with insignificant space
// characters elided.
func compactJSON(value string) string {
	buffer := bytes.Buffer{}
	trimmedValue := strings.TrimSuffix(value, "\n")
	json.Compact(&buffer, []byte(trimmedValue))

	return buffer.String()
}

func TestDecode(t *testing.T) {
	expectedOnlineMoverStruct := models.OnlineMover{
		Move:      "0adiC7Dr5WBppb01Mjub",
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC),
	}

	const onlineMoverJSON = `{
		"move": "0adiC7Dr5WBppb01Mjub",
		"mover": "5uls4pSbGeNvQFUYW8X74WraYcx2",
		"latitude": 43.481082,
		"longitude": -80.530143,
		"created_at": "2018-04-08T10:00:00Z"
	}`
	reader := strings.NewReader(onlineMoverJSON)
	actualOnlineMoverStruct := models.OnlineMover{}
	Decode(reader, &actualOnlineMoverStruct)

	assert.Equal(t, expectedOnlineMoverStruct, actualOnlineMoverStruct)
}

func TestEncode(t *testing.T) {
	expectedOnlineMoverJSON := `{
		"move": "0adiC7Dr5WBppb01Mjub",
		"mover": "5uls4pSbGeNvQFUYW8X74WraYcx2",
		"latitude": 43.481082,
		"longitude": -80.530143,
		"created_at": "2018-04-08T10:00:00Z"
	}`
	expectedOnlineMoverJSON = compactJSON(expectedOnlineMoverJSON)

	onlineMoverStruct := models.OnlineMover{
		Move:      "0adiC7Dr5WBppb01Mjub",
		Mover:     "5uls4pSbGeNvQFUYW8X74WraYcx2",
		Latitude:  43.481082,
		Longitude: -80.530143,
		CreatedAt: time.Date(2018, time.April, 8, 10, 0, 0, 0, time.UTC),
	}
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	Encode(writer, &onlineMoverStruct)
	writer.Flush()
	actualOnlineMoverJSON := compactJSON(buffer.String())

	assert.Equal(t, expectedOnlineMoverJSON, actualOnlineMoverJSON)
}
