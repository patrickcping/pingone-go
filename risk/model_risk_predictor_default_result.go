/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// RiskPredictorDefaultResult This specifies the result assigned to the predictor if the predictor could not be calculated during the risk evaluation. If this field is not provided, and the predictor could not be calculated during risk evaluation, the following options are: If the predictor is used in an override, the override is skipped. In the weighted policy, the predictor will have a weight of 0.
type RiskPredictorDefaultResult struct {
	// A string that identifies the risk level. Options are HIGH, MEDIUM, and LOW.
	Level string `json:"level"`
	Type *string `json:"type,omitempty"`
}

// NewRiskPredictorDefaultResult instantiates a new RiskPredictorDefaultResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorDefaultResult(level string) *RiskPredictorDefaultResult {
	this := RiskPredictorDefaultResult{}
	this.Level = level
	return &this
}

// NewRiskPredictorDefaultResultWithDefaults instantiates a new RiskPredictorDefaultResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorDefaultResultWithDefaults() *RiskPredictorDefaultResult {
	this := RiskPredictorDefaultResult{}
	return &this
}

// GetLevel returns the Level field value
func (o *RiskPredictorDefaultResult) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorDefaultResult) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *RiskPredictorDefaultResult) SetLevel(v string) {
	o.Level = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPredictorDefaultResult) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorDefaultResult) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPredictorDefaultResult) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskPredictorDefaultResult) SetType(v string) {
	o.Type = &v
}

func (o RiskPredictorDefaultResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPredictorDefaultResult struct {
	value *RiskPredictorDefaultResult
	isSet bool
}

func (v NullableRiskPredictorDefaultResult) Get() *RiskPredictorDefaultResult {
	return v.value
}

func (v *NullableRiskPredictorDefaultResult) Set(val *RiskPredictorDefaultResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorDefaultResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorDefaultResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorDefaultResult(val *RiskPredictorDefaultResult) *NullableRiskPredictorDefaultResult {
	return &NullableRiskPredictorDefaultResult{value: val, isSet: true}
}

func (v NullableRiskPredictorDefaultResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorDefaultResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


