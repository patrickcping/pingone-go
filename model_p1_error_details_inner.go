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

// P1ErrorDetailsInner struct for P1ErrorDetailsInner
type P1ErrorDetailsInner struct {
	Code *string `json:"code,omitempty"`
	Target *string `json:"target,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewP1ErrorDetailsInner instantiates a new P1ErrorDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP1ErrorDetailsInner() *P1ErrorDetailsInner {
	this := P1ErrorDetailsInner{}
	return &this
}

// NewP1ErrorDetailsInnerWithDefaults instantiates a new P1ErrorDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP1ErrorDetailsInnerWithDefaults() *P1ErrorDetailsInner {
	this := P1ErrorDetailsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *P1ErrorDetailsInner) SetCode(v string) {
	o.Code = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *P1ErrorDetailsInner) SetTarget(v string) {
	o.Target = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *P1ErrorDetailsInner) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInner) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *P1ErrorDetailsInner) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *P1ErrorDetailsInner) SetMessage(v string) {
	o.Message = &v
}

func (o P1ErrorDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableP1ErrorDetailsInner struct {
	value *P1ErrorDetailsInner
	isSet bool
}

func (v NullableP1ErrorDetailsInner) Get() *P1ErrorDetailsInner {
	return v.value
}

func (v *NullableP1ErrorDetailsInner) Set(val *P1ErrorDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableP1ErrorDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableP1ErrorDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP1ErrorDetailsInner(val *P1ErrorDetailsInner) *NullableP1ErrorDetailsInner {
	return &NullableP1ErrorDetailsInner{value: val, isSet: true}
}

func (v NullableP1ErrorDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP1ErrorDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

