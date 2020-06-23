package main

import "encoding/json"

func UnmarshalAction(data []byte) (Action, error) {
	var r Action
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Action) MarshalAction() ([]byte, error) {
	return json.Marshal(r)
}

type Action struct {
	Schema        string      `json:"schema"`
	Type          string      `json:"type"`
	Configuration interface{} `json:"configuration"`
}
