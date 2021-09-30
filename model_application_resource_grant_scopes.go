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

// ApplicationResourceGrantScopes struct for ApplicationResourceGrantScopes
type ApplicationResourceGrantScopes struct {
	// id A array that specifies the IDs of the scopes associated with this grant. This is a required property.
	Id *string `json:"id,omitempty"`
}

// NewApplicationResourceGrantScopes instantiates a new ApplicationResourceGrantScopes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourceGrantScopes() *ApplicationResourceGrantScopes {
	this := ApplicationResourceGrantScopes{}
	return &this
}

// NewApplicationResourceGrantScopesWithDefaults instantiates a new ApplicationResourceGrantScopes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceGrantScopesWithDefaults() *ApplicationResourceGrantScopes {
	this := ApplicationResourceGrantScopes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationResourceGrantScopes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrantScopes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationResourceGrantScopes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationResourceGrantScopes) SetId(v string) {
	o.Id = &v
}

func (o ApplicationResourceGrantScopes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResourceGrantScopes struct {
	value *ApplicationResourceGrantScopes
	isSet bool
}

func (v NullableApplicationResourceGrantScopes) Get() *ApplicationResourceGrantScopes {
	return v.value
}

func (v *NullableApplicationResourceGrantScopes) Set(val *ApplicationResourceGrantScopes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourceGrantScopes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourceGrantScopes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourceGrantScopes(val *ApplicationResourceGrantScopes) *NullableApplicationResourceGrantScopes {
	return &NullableApplicationResourceGrantScopes{value: val, isSet: true}
}

func (v NullableApplicationResourceGrantScopes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourceGrantScopes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

