package api

import (
	"github.com/muka/go-bluetooth/bluez/profile/device"
)

type BeaconType string

const (
	BeaconTypeEddystone = "eddystone"
	BeaconTypeIBeacon   = "ibeacon"
)

type Beacon struct {
	iBeacon   BeaconIBeacon
	eddystone BeaconEddystone
	Type      BeaconType
	Device    *device.Device1
}

func NewBeacon(dev *device.Device1) (bool, Beacon, error) {
	b := Beacon{
		Device: dev,
	}
	return b.Load(), b, nil
}

// IsEddystone return if the type of beacon is eddystone
func (b *Beacon) IsEddystone() bool {
	return b.Type == BeaconTypeEddystone
}

// IsIBeacon return if the type of beacon is ibeacon
func (b *Beacon) IsIBeacon() bool {
	return b.Type == BeaconTypeIBeacon
}

// GetEddystone return eddystone beacon information
func (b *Beacon) GetEddystone() BeaconEddystone {
	return b.eddystone
}

// GetIBeacon return if the type of beacon is ibeacon
func (b *Beacon) GetIBeacon() BeaconIBeacon {
	return b.iBeacon
}

// Load beacon inforamtion if available
func (b *Beacon) Load() bool {

	props := b.Device.Properties

	// log.Debugf("beacon props %++v", props)

	if len(props.ManufacturerData) > 0 {
		if frames, ok := props.ManufacturerData[0x76]; ok {
			// log.Debug("Found iBeacon")
			// log.Debugf("iBeacon data: %d", frames)
			b.Type = BeaconTypeIBeacon
			b.iBeacon = b.ParseIBeacon(frames.([]byte))
		}
		return true
	}

	for _, uuid := range props.UUIDs {
		if uuid == "FEAA" {
			if data, ok := props.ServiceData["FEAA"]; ok {
				// log.Debug("Found Eddystone")
				b.Type = BeaconTypeEddystone
				// log.Debugf("Eddystone data: %d", data)
				b.eddystone = b.ParseEddystone(data.([]byte))
			}
			return true
		}
	}

	return false
}
