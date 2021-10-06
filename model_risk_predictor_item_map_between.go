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

// RiskPredictorItemMapBetween Minimum and maximum boundaries
type RiskPredictorItemMapBetween struct {
	MinScore *int32 `json:"minScore,omitempty"`
	MaxScore *int32 `json:"maxScore,omitempty"`
}

// NewRiskPredictorItemMapBetween instantiates a new RiskPredictorItemMapBetween object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorItemMapBetween() *RiskPredictorItemMapBetween {
	this := RiskPredictorItemMapBetween{}
	return &this
}

// NewRiskPredictorItemMapBetweenWithDefaults instantiates a new RiskPredictorItemMapBetween object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorItemMapBetweenWithDefaults() *RiskPredictorItemMapBetween {
	this := RiskPredictorItemMapBetween{}
	return &this
}

// GetMinScore returns the MinScore field value if set, zero value otherwise.
func (o *RiskPredictorItemMapBetween) GetMinScore() int32 {
	if o == nil || o.MinScore == nil {
		var ret int32
		return ret
	}
	return *o.MinScore
}

// GetMinScoreOk returns a tuple with the MinScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMapBetween) GetMinScoreOk() (*int32, bool) {
	if o == nil || o.MinScore == nil {
		return nil, false
	}
	return o.MinScore, true
}

// HasMinScore returns a boolean if a field has been set.
func (o *RiskPredictorItemMapBetween) HasMinScore() bool {
	if o != nil && o.MinScore != nil {
		return true
	}

	return false
}

// SetMinScore gets a reference to the given int32 and assigns it to the MinScore field.
func (o *RiskPredictorItemMapBetween) SetMinScore(v int32) {
	o.MinScore = &v
}

// GetMaxScore returns the MaxScore field value if set, zero value otherwise.
func (o *RiskPredictorItemMapBetween) GetMaxScore() int32 {
	if o == nil || o.MaxScore == nil {
		var ret int32
		return ret
	}
	return *o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMapBetween) GetMaxScoreOk() (*int32, bool) {
	if o == nil || o.MaxScore == nil {
		return nil, false
	}
	return o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *RiskPredictorItemMapBetween) HasMaxScore() bool {
	if o != nil && o.MaxScore != nil {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given int32 and assigns it to the MaxScore field.
func (o *RiskPredictorItemMapBetween) SetMaxScore(v int32) {
	o.MaxScore = &v
}

func (o RiskPredictorItemMapBetween) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinScore != nil {
		toSerialize["minScore"] = o.MinScore
	}
	if o.MaxScore != nil {
		toSerialize["maxScore"] = o.MaxScore
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPredictorItemMapBetween struct {
	value *RiskPredictorItemMapBetween
	isSet bool
}

func (v NullableRiskPredictorItemMapBetween) Get() *RiskPredictorItemMapBetween {
	return v.value
}

func (v *NullableRiskPredictorItemMapBetween) Set(val *RiskPredictorItemMapBetween) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorItemMapBetween) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorItemMapBetween) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorItemMapBetween(val *RiskPredictorItemMapBetween) *NullableRiskPredictorItemMapBetween {
	return &NullableRiskPredictorItemMapBetween{value: val, isSet: true}
}

func (v NullableRiskPredictorItemMapBetween) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorItemMapBetween) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

