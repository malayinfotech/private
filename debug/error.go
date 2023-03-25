// Copyright (C) 2019 Storx Labs, Inc.
// See LICENSE for copying information.

package debug

import (
	"fmt"

	"github.com/spacemonkeygo/monkit/v3"

	"drpc/drpcerr"
)

func init() {
	monkit.AddErrorNameHandler(func(err error) (string, bool) {
		if code := drpcerr.Code(err); code != 0 {
			return fmt.Sprintf("drpc_%d", code), true
		}
		return "", false
	})
}
