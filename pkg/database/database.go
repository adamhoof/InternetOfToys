package database

type Device struct {
	MacAddress string
	Functions  []string
	// add other fields as needed
}

type Database interface {
	// GetDevice retrieves a device by its MAC address.
	GetDevice(macAddress string) (Device, error)

	// SaveDevice stores a new device in the database.
	SaveDevice(device Device) error

	// UpdateDevice updates an existing device in the database.
	UpdateDevice(device Device) error

	// DeleteDevice removes a device from the database.
	DeleteDevice(macAddress string) error

	// GetAllDevices retrieves all devices from the database.
	GetAllDevices() ([]Device, error)
}
