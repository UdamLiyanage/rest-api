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

func (r *DeviceSchema) MarshalDeviceSchema() ([]byte, error) {
	return json.Marshal(r)
}

type DeviceSchema struct {
	ID                    string        `json:"_id"`
	Version               string        `json:"version"`
	Name                  string        `json:"name"`
	Type                  string        `json:"type"`
	Public                string        `json:"public"`
	User                  string        `json:"user"`
	Enterprise            string        `json:"enterprise"`
	Communication         Communication `json:"communication"`
	Commands              []Command     `json:"commands"`
	Parameters            []Parameter   `json:"parameters"`
	Description           string        `json:"description"`
	LatestFirmwareVersion string        `json:"latest_firmware_version"`
	CommandFormat         CommandFormat `json:"command_format"`
	CreatedAt             time.Time     `json:"createdAt"`
	UpdatedAt             time.Time     `json:"updatedAt"`
}

type CommandFormat struct {
	FormatType string  `json:"format_type"`
	Payload    Payload `json:"payload"`
}

type Payload struct {
	Temp  string `json:"temp"`
	Power string `json:"power"`
}

type Command struct {
	Urn       string    `json:"urn"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
	Urn       string    `json:"urn"`
	Name      string    `json:"name"`
	Parameter string    `json:"parameter"`
	DataType  string    `json:"dataType"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
