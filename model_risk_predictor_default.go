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

// RiskPredictorDefault struct for RiskPredictorDefault
type RiskPredictorDefault struct {
	// An integer type. This specifies the weight assigned to the risk predictor in a new policy by default.
	Weight *int32 `json:"weight,omitempty"`
	Result *RiskPredictorDefaultResult `json:"result,omitempty"`
}

// NewRiskPredictorDefault instantiates a new RiskPredictorDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorDefault() *RiskPredictorDefault {
	this := RiskPredictorDefault{}
	return &this
}

// NewRiskPredictorDefaultWithDefaults instantiates a new RiskPredictorDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorDefaultWithDefaults() *RiskPredictorDefault {
	this := RiskPredictorDefault{}
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *RiskPredictorDefault) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorDefault) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *RiskPredictorDefault) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *RiskPredictorDefault) SetWeight(v int32) {
	o.Weight = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RiskPredictorDefault) GetResult() RiskPredictorDefaultResult {
	if o == nil || o.Result == nil {
		var ret RiskPredictorDefaultResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorDefault) GetResultOk() (*RiskPredictorDefaultResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RiskPredictorDefault) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given RiskPredictorDefaultResult and assigns it to the Result field.
func (o *RiskPredictorDefault) SetResult(v RiskPredictorDefaultResult) {
	o.Result = &v
}

func (o RiskPredictorDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPredictorDefault struct {
	value *RiskPredictorDefault
	isSet bool
}

func (v NullableRiskPredictorDefault) Get() *RiskPredictorDefault {
	return v.value
}

func (v *NullableRiskPredictorDefault) Set(val *RiskPredictorDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorDefault(val *RiskPredictorDefault) *NullableRiskPredictorDefault {
	return &NullableRiskPredictorDefault{value: val, isSet: true}
}

func (v NullableRiskPredictorDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


