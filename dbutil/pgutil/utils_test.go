// Copyright (C) 2019 Storx Labs, Inc.
// See LICENSE for copying information.

package pgutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"common/storx"
	"common/testrand"
)

func TestPostgresNodeIDsArray(t *testing.T) {
	ids := make(storx.NodeIDList, 10)
	for i := range ids {
		ids[i] = testrand.NodeID()
	}

	got, err := NodeIDArray(ids).Value() // returns a string
	require.NoError(t, err)

	expected, err := ByteaArray(ids.Bytes()).Value() // returns a string
	require.NoError(t, err)

	assert.Equal(t, expected.(string), got.(string))
}
