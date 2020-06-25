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
	ID                    string        `json:"_id,omitempty" bson:"_id"`
	Version               string        `json:"version" bson:"version"`
	Name                  string        `json:"name" bson:"name"`
	Type                  string        `json:"type" bson:"type"`
	Public                string        `json:"public"bson:"public"`
	User                  string        `json:"user" bson:"user"`
	Enterprise            string        `json:"enterprise" bson:"enterprise"`
	Communication         Communication `json:"communication" bson:"communication"`
	Commands              []Command     `json:"commands" bson:"commands"`
	Parameters            []Parameter   `json:"parameters" bson:"parameters"`
	Description           string        `json:"description" bson:"description"`
	LatestFirmwareVersion string        `json:"latest_firmware_version" bson:"latest_firmware_version"`
	CommandFormat         CommandFormat `json:"command_format" bson:"command_format"`
	CreatedAt             time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time     `json:"updatedAt" bson:"updatedAt"`
}

type CommandFormat struct {
	FormatType string  `json:"format_type" bson:"format_type"`
	Payload    Payload `json:"payload" bson:"payload"`
}

type Payload struct {
	Temp  string `json:"temp" bson:"temp"`
	Power string `json:"power" bson:"power"`
}

type Command struct {
	Urn       string    `json:"urn" bson:"urn"`
	Name      string    `json:"name" bson:"name"`
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Communication struct {
	Protocol    string      `json:"protocol" bson:"protocol"`
	Credentials Credentials `json:"credentials" bson:"credentials"`
}

type Credentials struct {
	URL            string         `json:"url" bson:"url"`
	Authentication Authentication `json:"authentication" bson:"authentication"`
}

type Authentication struct {
	Type     string `json:"type" bson:"type"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Parameter struct {
	Urn       string    `json:"urn" bson:"urn"`
	Name      string    `json:"name" bson:"name"`
	Parameter string    `json:"parameter" bson:"parameter"`
	DataType  string    `json:"dataType" bson:"dataType"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
