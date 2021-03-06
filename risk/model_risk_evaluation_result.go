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

// RiskEvaluationResult struct for RiskEvaluationResult
type RiskEvaluationResult struct {
	// A string that specifies the risk evaluation result type. Options are VALUE.
	Type *string `json:"type,omitempty"`
	// A string that specifies the risk evaluation result level. Options are HIGH, MEDIUM, and LOW.
	Level *string `json:"level,omitempty"`
	// A string that specifies any custom attribute the administrator defines.
	Value *string `json:"value,omitempty"`
}

// NewRiskEvaluationResult instantiates a new RiskEvaluationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationResult() *RiskEvaluationResult {
	this := RiskEvaluationResult{}
	return &this
}

// NewRiskEvaluationResultWithDefaults instantiates a new RiskEvaluationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationResultWithDefaults() *RiskEvaluationResult {
	this := RiskEvaluationResult{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskEvaluationResult) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationResult) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskEvaluationResult) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskEvaluationResult) SetType(v string) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RiskEvaluationResult) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationResult) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RiskEvaluationResult) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *RiskEvaluationResult) SetLevel(v string) {
	o.Level = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskEvaluationResult) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationResult) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskEvaluationResult) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskEvaluationResult) SetValue(v string) {
	o.Value = &v
}

func (o RiskEvaluationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationResult struct {
	value *RiskEvaluationResult
	isSet bool
}

func (v NullableRiskEvaluationResult) Get() *RiskEvaluationResult {
	return v.value
}

func (v *NullableRiskEvaluationResult) Set(val *RiskEvaluationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationResult(val *RiskEvaluationResult) *NullableRiskEvaluationResult {
	return &NullableRiskEvaluationResult{value: val, isSet: true}
}

func (v NullableRiskEvaluationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


