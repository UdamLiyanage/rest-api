package main

import (
	"encoding/json"
	"time"
)

func UnmarshalDeviceSchema(data []byte) (DeviceSchema, error) {
	var r DeviceSchema
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeviceSchema) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeviceSchema struct {
	ID                    string        `json:"_id"`
	Urn                   string        `json:"urn"`
	User                  string        `json:"user"`
	Version               string        `json:"version"`
	Name                  string        `json:"name"`
	Type                  string        `json:"type"`
	Public                string        `json:"public"`
	Communication         Communication `json:"communication"`
	Description           string        `json:"description"`
	Uid                   string        `json:"uid"`
	LatestFirmwareVersion string        `json:"latest_firmware_version"`
	Commands              []Command     `json:"commands"`
	Parameters            []Parameter   `json:"parameters"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
}

type Command struct {
	Name        string                 `json:"name"`
	DisplayName string                 `json:"display_name"`
	Type        string                 `json:"type"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}

type Communication struct {
	Protocol    string      `json:"protocol"`
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	URL            string         `json:"url"`
	Authentication Authentication `json:"authentication"`
}

type Authentication struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Parameter struct {
	Name      string `json:"name"`
	Parameter string `json:"parameter"`
	DataType  string `json:"data_type"`
}
