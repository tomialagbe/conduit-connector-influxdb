package influxdb

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

// version is set during the build process with ldflags (see Makefile).
// Default version matches default from runtime/debug.
var version = "0.1.0"

// Specification returns the connector's specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:        "influxdb",
		Summary:     "A Destination connector for InfluxDB.",
		Description: "A Destination connector for InfluxDB.",
		Version:     version,
		Author:      "Tomi Alagbe",
	}
}
