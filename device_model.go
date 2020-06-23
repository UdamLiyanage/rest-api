package main

import (
	"encoding/json"
	"time"
)

func UnmarshalDevice(data []byte) (Device, error) {
	var r Device
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Device) MarshalDevice() ([]byte, error) {
	return json.Marshal(r)
}

type Device struct {
	ID                 string    `json:"_id"`
	Serial             string    `json:"serial"`
	Urn                string    `json:"urn"`
	Schema             string    `json:"schema"`
	Name               string    `json:"name"`
	FirmwareVersion    string    `json:"firmware_version"`
	LastFirmwareUpdate string    `json:"last_firmware_update"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
