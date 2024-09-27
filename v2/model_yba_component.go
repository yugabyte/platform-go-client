/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"fmt"
)

// YbaComponent Individual YBA components to be backed up 
type YbaComponent string

// List of YbaComponent
const (
	YBA YbaComponent = "YBA"
	PROMETHEUS YbaComponent = "PROMETHEUS"
)

var allowedYbaComponentEnumValues = []YbaComponent{
	"YBA",
	"PROMETHEUS",
}

func (v *YbaComponent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := YbaComponent(value)
	for _, existing := range allowedYbaComponentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid YbaComponent", value)
}

// NewYbaComponentFromValue returns a pointer to a valid YbaComponent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewYbaComponentFromValue(v string) (*YbaComponent, error) {
	ev := YbaComponent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for YbaComponent: valid values are %v", v, allowedYbaComponentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v YbaComponent) IsValid() bool {
	for _, existing := range allowedYbaComponentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to YbaComponent value
func (v YbaComponent) Ptr() *YbaComponent {
	return &v
}

type NullableYbaComponent struct {
	value *YbaComponent
	isSet bool
}

func (v NullableYbaComponent) Get() *YbaComponent {
	return v.value
}

func (v *NullableYbaComponent) Set(val *YbaComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableYbaComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableYbaComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYbaComponent(val *YbaComponent) *NullableYbaComponent {
	return &NullableYbaComponent{value: val, isSet: true}
}

func (v NullableYbaComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYbaComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

