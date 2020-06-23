package main

import "encoding/json"

func UnmarshalRule(data []byte) (Rule, error) {
	var r Rule
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Rule) MarshalRule() ([]byte, error) {
	return json.Marshal(r)
}

type Rule struct {
	Device  string    `json:"device"`
	Actions []string  `json:"actions"`
	Rule    RuleClass `json:"rule"`
}

type RuleClass struct {
	Parameter string      `json:"parameter"`
	Condition string      `json:"condition"`
	Value     interface{} `json:"value"`
}
