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

// RoleAssignmentScope struct for RoleAssignmentScope
type RoleAssignmentScope struct {
	// A string that specifies the role assignment scope ID.
	Id string `json:"id"`
	// A string that specifies the type of resource defining the scope of the Role assignment. Options are ORGANIZATION, ENVIRONMENT, and POPULATION.
	Type string `json:"type"`
}

// NewRoleAssignmentScope instantiates a new RoleAssignmentScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentScope(id string, type_ string) *RoleAssignmentScope {
	this := RoleAssignmentScope{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewRoleAssignmentScopeWithDefaults instantiates a new RoleAssignmentScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentScopeWithDefaults() *RoleAssignmentScope {
	this := RoleAssignmentScope{}
	return &this
}

// GetId returns the Id field value
func (o *RoleAssignmentScope) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentScope) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleAssignmentScope) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RoleAssignmentScope) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentScope) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoleAssignmentScope) SetType(v string) {
	o.Type = v
}

func (o RoleAssignmentScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignmentScope struct {
	value *RoleAssignmentScope
	isSet bool
}

func (v NullableRoleAssignmentScope) Get() *RoleAssignmentScope {
	return v.value
}

func (v *NullableRoleAssignmentScope) Set(val *RoleAssignmentScope) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentScope) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentScope(val *RoleAssignmentScope) *NullableRoleAssignmentScope {
	return &NullableRoleAssignmentScope{value: val, isSet: true}
}

func (v NullableRoleAssignmentScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


