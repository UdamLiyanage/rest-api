package main

import "encoding/json"

func UnmarshalDeviceOwnership(data []byte) (DeviceOwnership, error) {
	var r DeviceOwnership
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeviceOwnership) MarshalDeviceOwnership() ([]byte, error) {
	return json.Marshal(r)
}

type DeviceOwnership struct {
	DeviceUrn string `json:"device_urn"`
	OwnerUrn  string `json:"owner_urn"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
