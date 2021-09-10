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

// RoleAssignmentEnvironment struct for RoleAssignmentEnvironment
type RoleAssignmentEnvironment struct {
	// A string that specifies the environment associated with the user.
	Id *string `json:"id,omitempty"`
}

// NewRoleAssignmentEnvironment instantiates a new RoleAssignmentEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentEnvironment() *RoleAssignmentEnvironment {
	this := RoleAssignmentEnvironment{}
	return &this
}

// NewRoleAssignmentEnvironmentWithDefaults instantiates a new RoleAssignmentEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentEnvironmentWithDefaults() *RoleAssignmentEnvironment {
	this := RoleAssignmentEnvironment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAssignmentEnvironment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentEnvironment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAssignmentEnvironment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleAssignmentEnvironment) SetId(v string) {
	o.Id = &v
}

func (o RoleAssignmentEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignmentEnvironment struct {
	value *RoleAssignmentEnvironment
	isSet bool
}

func (v NullableRoleAssignmentEnvironment) Get() *RoleAssignmentEnvironment {
	return v.value
}

func (v *NullableRoleAssignmentEnvironment) Set(val *RoleAssignmentEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentEnvironment(val *RoleAssignmentEnvironment) *NullableRoleAssignmentEnvironment {
	return &NullableRoleAssignmentEnvironment{value: val, isSet: true}
}

func (v NullableRoleAssignmentEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


