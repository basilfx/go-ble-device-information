package deviceinformation

import "github.com/go-ble/ble"

// ServiceDeviceInformation is de UUID of the Device Information service.
// https://www.bluetooth.com/specifications/gatt/services/.
var ServiceDeviceInformation = ble.UUID16(0x180a)

// UUIDs of the characteristics.
var (
	CharacteristicModelNumber      = ble.UUID16(0x2a24)
	CharacteristicSerialNumber     = ble.UUID16(0x2a25)
	CharacteristicFirmwareRevision = ble.UUID16(0x2a26)
	CharacteristicHardwareRevision = ble.UUID16(0x2a27)
	CharacteristicSoftwareRevision = ble.UUID16(0x2a28)
	CharacteristicManufacturer     = ble.UUID16(0x2a29)
)
