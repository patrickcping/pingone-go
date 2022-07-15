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

// SignOnPolicyActionMFABoundBiometrics Specifies MFA through a FIDO2 biometrics platform device.
type SignOnPolicyActionMFABoundBiometrics struct {
	// A boolean that specifies whether to permit users to authenticate with a FIDO2 biometrics platform device.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFABoundBiometrics instantiates a new SignOnPolicyActionMFABoundBiometrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFABoundBiometrics() *SignOnPolicyActionMFABoundBiometrics {
	this := SignOnPolicyActionMFABoundBiometrics{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFABoundBiometricsWithDefaults instantiates a new SignOnPolicyActionMFABoundBiometrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFABoundBiometricsWithDefaults() *SignOnPolicyActionMFABoundBiometrics {
	this := SignOnPolicyActionMFABoundBiometrics{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFABoundBiometrics) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFABoundBiometrics) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFABoundBiometrics) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFABoundBiometrics) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFABoundBiometrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFABoundBiometrics struct {
	value *SignOnPolicyActionMFABoundBiometrics
	isSet bool
}

func (v NullableSignOnPolicyActionMFABoundBiometrics) Get() *SignOnPolicyActionMFABoundBiometrics {
	return v.value
}

func (v *NullableSignOnPolicyActionMFABoundBiometrics) Set(val *SignOnPolicyActionMFABoundBiometrics) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFABoundBiometrics) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFABoundBiometrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFABoundBiometrics(val *SignOnPolicyActionMFABoundBiometrics) *NullableSignOnPolicyActionMFABoundBiometrics {
	return &NullableSignOnPolicyActionMFABoundBiometrics{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFABoundBiometrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFABoundBiometrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


