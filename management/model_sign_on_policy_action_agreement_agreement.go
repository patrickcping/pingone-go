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

// SignOnPolicyActionAgreementAgreement The relationship to the agreement to which the user must consent. The agreement must exist and be enabled. An agreement cannot be disabed if an action uses it. An enabled agreement must always support the default language. This property is required.
type SignOnPolicyActionAgreementAgreement struct {
	// A string that specifies the ID of the agreement to which the user must consent. This property is required.
	Id string `json:"id"`
}

// NewSignOnPolicyActionAgreementAgreement instantiates a new SignOnPolicyActionAgreementAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionAgreementAgreement(id string) *SignOnPolicyActionAgreementAgreement {
	this := SignOnPolicyActionAgreementAgreement{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionAgreementAgreementWithDefaults instantiates a new SignOnPolicyActionAgreementAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionAgreementAgreementWithDefaults() *SignOnPolicyActionAgreementAgreement {
	this := SignOnPolicyActionAgreementAgreement{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionAgreementAgreement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionAgreementAgreement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionAgreementAgreement) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionAgreementAgreement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionAgreementAgreement struct {
	value *SignOnPolicyActionAgreementAgreement
	isSet bool
}

func (v NullableSignOnPolicyActionAgreementAgreement) Get() *SignOnPolicyActionAgreementAgreement {
	return v.value
}

func (v *NullableSignOnPolicyActionAgreementAgreement) Set(val *SignOnPolicyActionAgreementAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionAgreementAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionAgreementAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionAgreementAgreement(val *SignOnPolicyActionAgreementAgreement) *NullableSignOnPolicyActionAgreementAgreement {
	return &NullableSignOnPolicyActionAgreementAgreement{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionAgreementAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionAgreementAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

