package influxdb

//go:generate paramgen -output=paramgen_dest.go DestinationConfig

import (
	"context"
	"fmt"

	influxdb3 "github.com/InfluxCommunity/influxdb3-go/influxdb3"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Destination struct {
	sdk.UnimplementedDestination
	config DestinationConfig
	client *influxdb3.Client
}

type DestinationConfig struct {
	// Config includes parameters that are the same in the source and destination.
	Config
	// DestinationConfigParam must be either yes or no (defaults to yes).
	DestinationConfigParam string `validate:"inclusion=yes|no" default:"yes"`
}

func NewDestination() sdk.Destination {
	// Create Destination and wrap it in the default middleware.
	return sdk.DestinationWithMiddleware(&Destination{}, sdk.DefaultDestinationMiddleware()...)
}

func (d *Destination) Parameters() map[string]sdk.Parameter {
	// Parameters is a map of named Parameters that describe how to configure
	// the Destination. Parameters can be generated from DestinationConfig with
	// paramgen.
	return d.config.Parameters()
}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
	// Configure is the first function to be called in a connector. It provides
	// the connector with the configuration that can be validated and stored.
	// In case the configuration is not valid it should return an error.
	// Testing if your connector can reach the configured data source should be
	// done in Open, not in Configure.
	// The SDK will validate the configuration and populate default values
	// before calling Configure. If you need to do more complex validations you
	// can do them manually here.

	sdk.Logger(ctx).Info().Msg("Configuring Destination...")
	err := sdk.Util.ParseConfig(cfg, &d.config)
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	return nil
}

func (d *Destination) Open(_ context.Context) error {
	// Open is called after Configure to signal the plugin it can prepare to
	// start writing records. If needed, the plugin should open connections in
	// this function.
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     d.config.InfluxDBHost,
		Token:    d.config.InfluxDBToken,
		Database: d.config.InfluxDBDatabase,
	})
	if err != nil {
		return err
	}
	d.client = client
	return nil
}

func (d *Destination) Write(context context.Context, records []sdk.Record) (int, error) {
	// Write writes len(r) records from r to the destination right away without
	// caching. It should return the number of records written from r
	// (0 <= n <= len(r)) and any error encountered that caused the write to
	// stop early. Write must return a non-nil error if it returns n < len(r).
	for i, r := range records {
		fmt.Printf("Record %d: %v\n", i, r)
		err := d.client.Write(context, r.Payload.After.Bytes())
		if err != nil {
			return i, fmt.Errorf("failed to write record to influx record index: %v, error: %w", i, err)
		}
	}

	return len(records), nil
}

func (d *Destination) Teardown(_ context.Context) error {
	// Teardown signals to the plugin that all records were written and there
	// will be no more calls to any other function. After Teardown returns, the
	// plugin should be ready for a graceful shutdown.
	if d.client != nil {
		return d.client.Close()
	}
	return nil
}
