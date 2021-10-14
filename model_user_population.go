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

// UserPopulation struct for UserPopulation
type UserPopulation struct {
	// A string that specifies the identifier of the population resource associated with the user. This property cannot be updated using PUT {{apiPath}}/environments/{{envID}}/users/{{userID}}. However, it can be updated using PUT /environments/{{envID}}/users/{{userID}}/population.
	Id string `json:"id"`
}

// NewUserPopulation instantiates a new UserPopulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPopulation(id string) *UserPopulation {
	this := UserPopulation{}
	this.Id = id
	return &this
}

// NewUserPopulationWithDefaults instantiates a new UserPopulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPopulationWithDefaults() *UserPopulation {
	this := UserPopulation{}
	return &this
}

// GetId returns the Id field value
func (o *UserPopulation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserPopulation) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserPopulation) SetId(v string) {
	o.Id = v
}

func (o UserPopulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUserPopulation struct {
	value *UserPopulation
	isSet bool
}

func (v NullableUserPopulation) Get() *UserPopulation {
	return v.value
}

func (v *NullableUserPopulation) Set(val *UserPopulation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPopulation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPopulation(val *UserPopulation) *NullableUserPopulation {
	return &NullableUserPopulation{value: val, isSet: true}
}

func (v NullableUserPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


