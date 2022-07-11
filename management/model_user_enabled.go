/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// UserEnabled struct for UserEnabled
type UserEnabled struct {
	// A read-only boolean attribute that specifies whether the user is enabled. This attribute is set to ‘true’ by default when the user is created.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUserEnabled instantiates a new UserEnabled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEnabled() *UserEnabled {
	this := UserEnabled{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewUserEnabledWithDefaults instantiates a new UserEnabled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEnabledWithDefaults() *UserEnabled {
	this := UserEnabled{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserEnabled) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEnabled) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserEnabled) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserEnabled) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UserEnabled) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableUserEnabled struct {
	value *UserEnabled
	isSet bool
}

func (v NullableUserEnabled) Get() *UserEnabled {
	return v.value
}

func (v *NullableUserEnabled) Set(val *UserEnabled) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEnabled) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEnabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEnabled(val *UserEnabled) *NullableUserEnabled {
	return &NullableUserEnabled{value: val, isSet: true}
}

func (v NullableUserEnabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEnabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


