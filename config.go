// package connectorname

// // Config contains shared config parameters, common to the source and
// // destination. If you don't need shared parameters you can entirely remove this
// // file.
// type Config struct {
// 	// GlobalConfigParam is named global_config_param_name and needs to be
// 	// provided by the user.
// 	GlobalConfigParam string `json:"global_config_param_name" validate:"required"`
// }

package influxdb

// Config contains shared config parameters, common to the source and
// destination. If you don't need shared parameters you can entirely remove this
// file.
type Config struct {
	// InfluxDBHost is named INFLUXDB_HOST and needs to be
	// provided by the user.
	InfluxDBHost string `json:"INFLUXDB_HOST" validate:"required"`
	// InfluxDBToken is named INFLUXDB_TOKEN and needs to be
	// provided by the user.
	InfluxDBToken string `json:"INFLUXDB_TOKEN" validate:"required"`

	// InfluxDBDatabase is named INFLUXDB_DATABASE and needs to be
	// provided by the user.
	InfluxDBDatabase string `json:"INFLUXDB_DATABASE" validate:"required"`
}
