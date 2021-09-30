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

// RiskEvaluationEventEvaluatedFactors struct for RiskEvaluationEventEvaluatedFactors
type RiskEvaluationEventEvaluatedFactors struct {
	// A string that specifies the state of the transaction. Options are FAILED, IN_PROGRESS, and SUCCESS.
	Status *string `json:"status,omitempty"`
	// A string that specifies the transaction type.
	Type *string `json:"type,omitempty"`
}

// NewRiskEvaluationEventEvaluatedFactors instantiates a new RiskEvaluationEventEvaluatedFactors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationEventEvaluatedFactors() *RiskEvaluationEventEvaluatedFactors {
	this := RiskEvaluationEventEvaluatedFactors{}
	return &this
}

// NewRiskEvaluationEventEvaluatedFactorsWithDefaults instantiates a new RiskEvaluationEventEvaluatedFactors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationEventEvaluatedFactorsWithDefaults() *RiskEvaluationEventEvaluatedFactors {
	this := RiskEvaluationEventEvaluatedFactors{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RiskEvaluationEventEvaluatedFactors) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventEvaluatedFactors) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RiskEvaluationEventEvaluatedFactors) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RiskEvaluationEventEvaluatedFactors) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskEvaluationEventEvaluatedFactors) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventEvaluatedFactors) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskEvaluationEventEvaluatedFactors) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskEvaluationEventEvaluatedFactors) SetType(v string) {
	o.Type = &v
}

func (o RiskEvaluationEventEvaluatedFactors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationEventEvaluatedFactors struct {
	value *RiskEvaluationEventEvaluatedFactors
	isSet bool
}

func (v NullableRiskEvaluationEventEvaluatedFactors) Get() *RiskEvaluationEventEvaluatedFactors {
	return v.value
}

func (v *NullableRiskEvaluationEventEvaluatedFactors) Set(val *RiskEvaluationEventEvaluatedFactors) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationEventEvaluatedFactors) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationEventEvaluatedFactors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationEventEvaluatedFactors(val *RiskEvaluationEventEvaluatedFactors) *NullableRiskEvaluationEventEvaluatedFactors {
	return &NullableRiskEvaluationEventEvaluatedFactors{value: val, isSet: true}
}

func (v NullableRiskEvaluationEventEvaluatedFactors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationEventEvaluatedFactors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

