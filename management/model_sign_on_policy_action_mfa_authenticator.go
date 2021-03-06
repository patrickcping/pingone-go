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

// SignOnPolicyActionMFAAuthenticator Specifies MFA through a Time-based One-time Password (TOTP) option over a third-party authenticator application.
type SignOnPolicyActionMFAAuthenticator struct {
	// A boolean that specifies the enabled/disabled state of the MFA through a Time-based One-time Password (TOTP) action. The default is disabled when creating a new policy. When enabled, it allows users to authenticate using a TOTP authenticator application.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFAAuthenticator instantiates a new SignOnPolicyActionMFAAuthenticator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAAuthenticator() *SignOnPolicyActionMFAAuthenticator {
	this := SignOnPolicyActionMFAAuthenticator{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFAAuthenticatorWithDefaults instantiates a new SignOnPolicyActionMFAAuthenticator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAAuthenticatorWithDefaults() *SignOnPolicyActionMFAAuthenticator {
	this := SignOnPolicyActionMFAAuthenticator{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFAAuthenticator) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAuthenticator) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFAAuthenticator) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFAAuthenticator) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFAAuthenticator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFAAuthenticator struct {
	value *SignOnPolicyActionMFAAuthenticator
	isSet bool
}

func (v NullableSignOnPolicyActionMFAAuthenticator) Get() *SignOnPolicyActionMFAAuthenticator {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAAuthenticator) Set(val *SignOnPolicyActionMFAAuthenticator) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAAuthenticator) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAAuthenticator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAAuthenticator(val *SignOnPolicyActionMFAAuthenticator) *NullableSignOnPolicyActionMFAAuthenticator {
	return &NullableSignOnPolicyActionMFAAuthenticator{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAAuthenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAAuthenticator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


