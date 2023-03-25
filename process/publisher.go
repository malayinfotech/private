// Copyright (C) 2022 Storx Labs, Inc.
// See LICENSE for copying information.

package process

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"

	"common/eventstat"
	"common/telemetry"
)

// InitEventStatPublisherWithHostname initializes telemetry reporting, using the hostname as the telemetry instance ID.
// ctx should be cancelled to stop the telemetry publisher.
func InitEventStatPublisherWithHostname(ctx context.Context, log *zap.Logger, r *eventstat.Registry) error {
	var metricsID string
	hostname, err := os.Hostname()
	if err != nil {
		log.Warn("Could not read hostname for telemetry setup", zap.Error(err))
		metricsID = telemetry.DefaultInstanceID()
	} else {
		metricsID = strings.ReplaceAll(hostname, ".", "_")
	}

	instanceID := *metricInstancePrefix + metricsID
	if len(instanceID) > maxInstanceLength {
		instanceID = instanceID[:maxInstanceLength]
	}

	return InitEventStatPublisher(ctx, log, r, func(opts *eventstat.ClientOpts) {
		opts.Instance = instanceID
	})
}

// InitEventStatPublisher initializes telemetry reporting.
func InitEventStatPublisher(ctx context.Context, log *zap.Logger, r *eventstat.Registry, customization func(*eventstat.ClientOpts)) error {
	collectors := strings.Split(*metricCollector, ",")
	if len(collectors) > 1 {
		log.Warn("Event stat can be published only to one collector server")
	}
	if len(collectors) == 1 {
		opts := &eventstat.ClientOpts{
			Interval: calcMetricInterval(),
		}
		customization(opts)
		publisher, err := eventstat.NewUDPPublisher(collectors[0], r, *opts)
		if err != nil {
			return err
		}

		go publisher.Run(ctx)
		log.Info("Event stat publisher is enabled", zap.String("instance ID", opts.Instance))
	}
	return nil
}
