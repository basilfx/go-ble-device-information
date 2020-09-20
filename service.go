package deviceinformation

import "github.com/go-ble/ble"

// Service is the interface for the device information BLE service. Only
// properties that are non-nil will be advertised by this service.
type Service struct {
	modelNumber      *string
	serialNumber     *string
	firmwareRevision *string
	hardwareRevision *string
	softwareRevision *string
	manufacturer     *string
}

// New initializes a new instance of Service. Only non-nil
// values are advertised.
func New() *Service {
	return &Service{}
}

// SetModelNumber sets the model number.
func (d *Service) SetModelNumber(modelNumber string) {
	d.modelNumber = &modelNumber
}

// SetSerialNumber sets the serial number.
func (d *Service) SetSerialNumber(serialNumber string) {
	d.serialNumber = &serialNumber
}

// SetFirmwareRevision sets the firmware revision.
func (d *Service) SetFirmwareRevision(firmwareRevision string) {
	d.firmwareRevision = &firmwareRevision
}

// SetHardwareRevision sets the hardware revision.
func (d *Service) SetHardwareRevision(hardwareRevision string) {
	d.hardwareRevision = &hardwareRevision
}

// SetSoftwareRevision sets the software revision.
func (d *Service) SetSoftwareRevision(softwareRevision string) {
	d.softwareRevision = &softwareRevision
}

// SetManufacturer sets the manufacturer.
func (d *Service) SetManufacturer(manufacturer string) {
	d.manufacturer = &manufacturer
}

// Create will return an instance of ble.Service, that can be used to advertise
// the device information service.
func (d *Service) Create() *ble.Service {
	s := ble.NewService(ServiceDeviceInformation)

	// Model Number.
	if d.modelNumber != nil {
		char := s.NewCharacteristic(CharacteristicModelNumber)

		char.SetValue([]byte(*d.modelNumber))
	}

	// Serial Number.
	if d.serialNumber != nil {
		char := s.NewCharacteristic(CharacteristicSerialNumber)

		char.SetValue([]byte(*d.serialNumber))
	}

	// Firmware Revision.
	if d.firmwareRevision != nil {
		char := s.NewCharacteristic(CharacteristicFirmwareRevision)

		char.SetValue([]byte(*d.firmwareRevision))
	}

	// Hardware Revision.
	if d.hardwareRevision != nil {
		char := s.NewCharacteristic(CharacteristicHardwareRevision)

		char.SetValue([]byte(*d.hardwareRevision))
	}

	// Software Revision.
	if d.softwareRevision != nil {
		char := s.NewCharacteristic(CharacteristicSoftwareRevision)

		char.SetValue([]byte(*d.softwareRevision))
	}

	// Manufacturer.
	if d.manufacturer != nil {
		char := s.NewCharacteristic(CharacteristicManufacturer)

		char.SetValue([]byte(*d.manufacturer))
	}

	return s
}
