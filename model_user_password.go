/*
PingOne Platform API - Management

A bare-bones collection for the PingOne API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"encoding/json"
)

// UserPassword An object that specifies the user's password. This property is never returned in the response.
type UserPassword struct {
	// (Optional) A boolean that specifies whether the user is forced to change the password on the next log in. If not provided, the property is set to false.
	ForceChange *bool `json:"forceChange,omitempty"`
	// A string that specifies the user's password value. The string is either in cleartext or pre-encoded format.
	Value *string `json:"value,omitempty"`
	External *UserPasswordExternal `json:"external,omitempty"`
}

// NewUserPassword instantiates a new UserPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPassword() *UserPassword {
	this := UserPassword{}
	return &this
}

// NewUserPasswordWithDefaults instantiates a new UserPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordWithDefaults() *UserPassword {
	this := UserPassword{}
	return &this
}

// GetForceChange returns the ForceChange field value if set, zero value otherwise.
func (o *UserPassword) GetForceChange() bool {
	if o == nil || o.ForceChange == nil {
		var ret bool
		return ret
	}
	return *o.ForceChange
}

// GetForceChangeOk returns a tuple with the ForceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassword) GetForceChangeOk() (*bool, bool) {
	if o == nil || o.ForceChange == nil {
		return nil, false
	}
	return o.ForceChange, true
}

// HasForceChange returns a boolean if a field has been set.
func (o *UserPassword) HasForceChange() bool {
	if o != nil && o.ForceChange != nil {
		return true
	}

	return false
}

// SetForceChange gets a reference to the given bool and assigns it to the ForceChange field.
func (o *UserPassword) SetForceChange(v bool) {
	o.ForceChange = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserPassword) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassword) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserPassword) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserPassword) SetValue(v string) {
	o.Value = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *UserPassword) GetExternal() UserPasswordExternal {
	if o == nil || o.External == nil {
		var ret UserPasswordExternal
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPassword) GetExternalOk() (*UserPasswordExternal, bool) {
	if o == nil || o.External == nil {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *UserPassword) HasExternal() bool {
	if o != nil && o.External != nil {
		return true
	}

	return false
}

// SetExternal gets a reference to the given UserPasswordExternal and assigns it to the External field.
func (o *UserPassword) SetExternal(v UserPasswordExternal) {
	o.External = &v
}

func (o UserPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceChange != nil {
		toSerialize["forceChange"] = o.ForceChange
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.External != nil {
		toSerialize["external"] = o.External
	}
	return json.Marshal(toSerialize)
}

type NullableUserPassword struct {
	value *UserPassword
	isSet bool
}

func (v NullableUserPassword) Get() *UserPassword {
	return v.value
}

func (v *NullableUserPassword) Set(val *UserPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPassword(val *UserPassword) *NullableUserPassword {
	return &NullableUserPassword{value: val, isSet: true}
}

func (v NullableUserPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


