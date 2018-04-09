/*
 * Copyright 2017-present, DoorTwoDoor, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the Apache-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package utilities provides functions to work with JSON codec, write
// responses and perform Cloud Datastore operations.
package utilities

import (
	"encoding/json"
	"io"
)

// Decode reads the next JSON-encoded value from its input and stores it in the
// value pointed to by value.
func Decode(reader io.Reader, value interface{}) error {
	return json.NewDecoder(reader).Decode(value)
}

// Encode writes the JSON encoding of v to the stream, followed by a newline
// character.
func Encode(writer io.Writer, value interface{}) error {
	return json.NewEncoder(writer).Encode(value)
}
