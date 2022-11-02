package humidity_temperature

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

type Device struct {
	dht.Device
}

// Setup configures a DHT22 sensor on pin p and polls every t interval,
// the smallest value t is allowed to take is 2 seconds
func Setup(p machine.Pin, t time.Duration) Device {
	// set sensor type
	sensor := Device{dht.New(p, dht.DHT22)}

	// configure update policy
	sensor.Configure(dht.UpdatePolicy{UpdateTime: t})

	return sensor
}

// GetTemperature receives a temperature measurement and returns it
// alongside any error that occurred
func (d *Device) GetTemperature() (int16, error) {
	// get measurements from bus
	d.ReadMeasurements()

	temp, err := d.Temperature()
	return temp, err
}

// GetHumidity receives a humidity measurement and returns it
// alongside any error that occurred
func (d *Device) GetHumidity() (uint16, error) {
	// get measurements from bus
	d.ReadMeasurements()

	hum, err := d.Humidity()
	return hum, err
}

// GetAll retrieves Temperature and Humidity readings and returns it
// alongside any error that occurred
func (d *Device) GetAll() (int16, uint16, error) {
	d.ReadMeasurements()

	temp, hum, err := d.Measurements()
	return temp, hum, err
}
