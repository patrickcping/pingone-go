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

// SignOnPolicyActionProgressiveProfilingAttributes struct for SignOnPolicyActionProgressiveProfilingAttributes
type SignOnPolicyActionProgressiveProfilingAttributes struct {
	// A string that specifies the name and path of the user profile attribute as defined in the user schema (for example, email or address.postalCode). This property is required.
	Name string `json:"name"`
	// A boolean that specifies whether the user is required to provide a value for the attribute. This property is required.
	Required bool `json:"required"`
}

// NewSignOnPolicyActionProgressiveProfilingAttributes instantiates a new SignOnPolicyActionProgressiveProfilingAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionProgressiveProfilingAttributes(name string, required bool) *SignOnPolicyActionProgressiveProfilingAttributes {
	this := SignOnPolicyActionProgressiveProfilingAttributes{}
	this.Name = name
	this.Required = required
	return &this
}

// NewSignOnPolicyActionProgressiveProfilingAttributesWithDefaults instantiates a new SignOnPolicyActionProgressiveProfilingAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionProgressiveProfilingAttributesWithDefaults() *SignOnPolicyActionProgressiveProfilingAttributes {
	this := SignOnPolicyActionProgressiveProfilingAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *SignOnPolicyActionProgressiveProfilingAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfilingAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignOnPolicyActionProgressiveProfilingAttributes) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value
func (o *SignOnPolicyActionProgressiveProfilingAttributes) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfilingAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *SignOnPolicyActionProgressiveProfilingAttributes) SetRequired(v bool) {
	o.Required = v
}

func (o SignOnPolicyActionProgressiveProfilingAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["required"] = o.Required
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionProgressiveProfilingAttributes struct {
	value *SignOnPolicyActionProgressiveProfilingAttributes
	isSet bool
}

func (v NullableSignOnPolicyActionProgressiveProfilingAttributes) Get() *SignOnPolicyActionProgressiveProfilingAttributes {
	return v.value
}

func (v *NullableSignOnPolicyActionProgressiveProfilingAttributes) Set(val *SignOnPolicyActionProgressiveProfilingAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionProgressiveProfilingAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionProgressiveProfilingAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionProgressiveProfilingAttributes(val *SignOnPolicyActionProgressiveProfilingAttributes) *NullableSignOnPolicyActionProgressiveProfilingAttributes {
	return &NullableSignOnPolicyActionProgressiveProfilingAttributes{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionProgressiveProfilingAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionProgressiveProfilingAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

