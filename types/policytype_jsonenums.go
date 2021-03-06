// generated by jsonenums -type=PolicyType; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

var (
	_PolicyTypeNameToValue = map[string]PolicyType{
		"PolicyTypeNone":  PolicyTypeNone,
		"PolicyTypeAll":   PolicyTypeAll,
		"PolicyTypeMajor": PolicyTypeMajor,
		"PolicyTypeMinor": PolicyTypeMinor,
		"PolicyTypePatch": PolicyTypePatch,
		"PolicyTypeForce": PolicyTypeForce,
	}

	_PolicyTypeValueToName = map[PolicyType]string{
		PolicyTypeNone:  "PolicyTypeNone",
		PolicyTypeAll:   "PolicyTypeAll",
		PolicyTypeMajor: "PolicyTypeMajor",
		PolicyTypeMinor: "PolicyTypeMinor",
		PolicyTypePatch: "PolicyTypePatch",
		PolicyTypeForce: "PolicyTypeForce",
	}
)

func init() {
	var v PolicyType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_PolicyTypeNameToValue = map[string]PolicyType{
			interface{}(PolicyTypeNone).(fmt.Stringer).String():  PolicyTypeNone,
			interface{}(PolicyTypeAll).(fmt.Stringer).String():   PolicyTypeAll,
			interface{}(PolicyTypeMajor).(fmt.Stringer).String(): PolicyTypeMajor,
			interface{}(PolicyTypeMinor).(fmt.Stringer).String(): PolicyTypeMinor,
			interface{}(PolicyTypePatch).(fmt.Stringer).String(): PolicyTypePatch,
			interface{}(PolicyTypeForce).(fmt.Stringer).String(): PolicyTypeForce,
		}
	}
}

// MarshalJSON is generated so PolicyType satisfies json.Marshaler.
func (r PolicyType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PolicyTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid PolicyType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so PolicyType satisfies json.Unmarshaler.
func (r *PolicyType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PolicyType should be a string, got %s", data)
	}
	v, ok := _PolicyTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PolicyType %q", s)
	}
	*r = v
	return nil
}
