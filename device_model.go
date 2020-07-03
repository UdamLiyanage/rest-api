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
	Serial             string    `json:"serial" bson:"serial"`
	Urn                string    `json:"urn" bson:"urn"`
	Schema             string    `json:"schema" bson:"schema"`
	User               string    `json:"user" bson:"user"`
	Enterprise         string    `json:"enterprise" bson:"enterprise"`
	Name               string    `json:"name" bson:"name"`
	FirmwareVersion    string    `json:"firmware_version" bson:"firmware_version"`
	LastFirmwareUpdate string    `json:"last_firmware_update" bson:"last_firmware_update"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}
