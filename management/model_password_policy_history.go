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

// PasswordPolicyHistory Settings to control the users password history
type PasswordPolicyHistory struct {
	// Specifies the number of prior passwords to keep for prevention of password re-use. The value must be a positive, non-zero integer.
	Count *int32 `json:"count,omitempty"`
	// The length of time to keep recent passwords for prevention of password re-use. The value must be a positive, non-zero integer.
	RetentionDays *int32 `json:"retentionDays,omitempty"`
}

// NewPasswordPolicyHistory instantiates a new PasswordPolicyHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyHistory() *PasswordPolicyHistory {
	this := PasswordPolicyHistory{}
	return &this
}

// NewPasswordPolicyHistoryWithDefaults instantiates a new PasswordPolicyHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyHistoryWithDefaults() *PasswordPolicyHistory {
	this := PasswordPolicyHistory{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PasswordPolicyHistory) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyHistory) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PasswordPolicyHistory) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PasswordPolicyHistory) SetCount(v int32) {
	o.Count = &v
}

// GetRetentionDays returns the RetentionDays field value if set, zero value otherwise.
func (o *PasswordPolicyHistory) GetRetentionDays() int32 {
	if o == nil || o.RetentionDays == nil {
		var ret int32
		return ret
	}
	return *o.RetentionDays
}

// GetRetentionDaysOk returns a tuple with the RetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyHistory) GetRetentionDaysOk() (*int32, bool) {
	if o == nil || o.RetentionDays == nil {
		return nil, false
	}
	return o.RetentionDays, true
}

// HasRetentionDays returns a boolean if a field has been set.
func (o *PasswordPolicyHistory) HasRetentionDays() bool {
	if o != nil && o.RetentionDays != nil {
		return true
	}

	return false
}

// SetRetentionDays gets a reference to the given int32 and assigns it to the RetentionDays field.
func (o *PasswordPolicyHistory) SetRetentionDays(v int32) {
	o.RetentionDays = &v
}

func (o PasswordPolicyHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.RetentionDays != nil {
		toSerialize["retentionDays"] = o.RetentionDays
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordPolicyHistory struct {
	value *PasswordPolicyHistory
	isSet bool
}

func (v NullablePasswordPolicyHistory) Get() *PasswordPolicyHistory {
	return v.value
}

func (v *NullablePasswordPolicyHistory) Set(val *PasswordPolicyHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyHistory(val *PasswordPolicyHistory) *NullablePasswordPolicyHistory {
	return &NullablePasswordPolicyHistory{value: val, isSet: true}
}

func (v NullablePasswordPolicyHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


