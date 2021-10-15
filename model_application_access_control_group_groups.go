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

// ApplicationAccessControlGroupGroups struct for ApplicationAccessControlGroupGroups
type ApplicationAccessControlGroupGroups struct {
	Id string `json:"id"`
}

// NewApplicationAccessControlGroupGroups instantiates a new ApplicationAccessControlGroupGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAccessControlGroupGroups(id string) *ApplicationAccessControlGroupGroups {
	this := ApplicationAccessControlGroupGroups{}
	this.Id = id
	return &this
}

// NewApplicationAccessControlGroupGroupsWithDefaults instantiates a new ApplicationAccessControlGroupGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAccessControlGroupGroupsWithDefaults() *ApplicationAccessControlGroupGroups {
	this := ApplicationAccessControlGroupGroups{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationAccessControlGroupGroups) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationAccessControlGroupGroups) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationAccessControlGroupGroups) SetId(v string) {
	o.Id = v
}

func (o ApplicationAccessControlGroupGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationAccessControlGroupGroups struct {
	value *ApplicationAccessControlGroupGroups
	isSet bool
}

func (v NullableApplicationAccessControlGroupGroups) Get() *ApplicationAccessControlGroupGroups {
	return v.value
}

func (v *NullableApplicationAccessControlGroupGroups) Set(val *ApplicationAccessControlGroupGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAccessControlGroupGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAccessControlGroupGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAccessControlGroupGroups(val *ApplicationAccessControlGroupGroups) *NullableApplicationAccessControlGroupGroups {
	return &NullableApplicationAccessControlGroupGroups{value: val, isSet: true}
}

func (v NullableApplicationAccessControlGroupGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAccessControlGroupGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

