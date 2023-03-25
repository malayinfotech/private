// Copyright (C) 2020 Storx Labs, Inc.
// See LICENSE for copying information.

package process

import (
	"go.uber.org/zap/zapcore"
)

func levelDecorate(level zapcore.Level, message string) string {
	return message
}
