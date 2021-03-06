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

// RiskPredictorItemMapBetween Minimum and maximum boundaries
type RiskPredictorItemMapBetween struct {
	MinScore float32 `json:"minScore"`
	MaxScore float32 `json:"maxScore"`
}

// NewRiskPredictorItemMapBetween instantiates a new RiskPredictorItemMapBetween object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorItemMapBetween(minScore float32, maxScore float32) *RiskPredictorItemMapBetween {
	this := RiskPredictorItemMapBetween{}
	this.MinScore = minScore
	this.MaxScore = maxScore
	return &this
}

// NewRiskPredictorItemMapBetweenWithDefaults instantiates a new RiskPredictorItemMapBetween object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorItemMapBetweenWithDefaults() *RiskPredictorItemMapBetween {
	this := RiskPredictorItemMapBetween{}
	return &this
}

// GetMinScore returns the MinScore field value
func (o *RiskPredictorItemMapBetween) GetMinScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinScore
}

// GetMinScoreOk returns a tuple with the MinScore field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMapBetween) GetMinScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinScore, true
}

// SetMinScore sets field value
func (o *RiskPredictorItemMapBetween) SetMinScore(v float32) {
	o.MinScore = v
}

// GetMaxScore returns the MaxScore field value
func (o *RiskPredictorItemMapBetween) GetMaxScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMapBetween) GetMaxScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxScore, true
}

// SetMaxScore sets field value
func (o *RiskPredictorItemMapBetween) SetMaxScore(v float32) {
	o.MaxScore = v
}

func (o RiskPredictorItemMapBetween) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["minScore"] = o.MinScore
	}
	if true {
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


