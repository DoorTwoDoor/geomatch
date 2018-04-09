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
	"context"

	"google.golang.org/appengine/datastore"
)

// PutToDatastore saves an entity into the datastore with an automatically
// generated key.
func PutToDatastore(
	context context.Context,
	kind string,
	value interface{},
) (*datastore.Key, error) {
	key := datastore.NewIncompleteKey(context, kind, nil)

	return datastore.Put(context, key, value)
}
