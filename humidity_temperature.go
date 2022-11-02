package humiditytemperature

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

// Setup configures a DHT22 sensor on pin p and polls every t interval.
// the smallest value t is allowed to take is 2 seconds
func Setup(p machine.Pin, t time.Duration) {
	// set sensor type
	sensor := dht.New(p, dht.DHT22)

	// configure update policy
	sensor.Configure(dht.UpdatePolicy{UpdateTime: t})
}

func GetTemperature() {}

func GetHumidity() {}
