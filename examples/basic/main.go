package main

import (
	"context"
	"fmt"

	deviceinformation "github.com/basilfx/go-ble-device-information"
	"github.com/basilfx/go-ble-utilities/device"

	"github.com/go-ble/ble"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Create a new device.
	d, err := device.NewDevice()

	if err != nil {
		log.Fatalf("Unable to create new device: %v", err)
	}

	ble.SetDefaultDevice(d)

	// Add device information service.
	service := deviceinformation.New()

	service.SetModelNumber("My Device")
	service.SetManufacturer("Example.org")
	service.SetFirmwareRevision("1.0.0")
	service.SetHardwareRevision("1.0.0")

	err = ble.AddService(service.Create())

	if err != nil {
		log.Fatalf("Unable to add device information service: %v", err)
	}

	// Advertise for specified durantion, or until interrupted by user.
	ctx := ble.WithSigHandler(context.WithCancel(context.Background()))

	err = ble.AdvertiseNameAndServices(ctx, "My Device")

	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}
