package main

import (
	influxdb "conduit-connector-influxdb"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(influxdb.Connector)
}
