// Copyright (C) 2022 Storx Labs, Inc.
// See LICENSE for copying information.

//go:build !go1.18
// +build !go1.18

package version

func init() {
	Build = getInfoFromBuildTags()
}
