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

// RiskEvaluationDetailsUserVelocityByIp struct for RiskEvaluationDetailsUserVelocityByIp
type RiskEvaluationDetailsUserVelocityByIp struct {
	// An enum indicating whether the calculated number of users per IP is LOW, MEDIUM, or HIGH.
	Level *string `json:"level,omitempty"`
	// A string indicating the reason the user was flagged. For example \"More than 250 users accessed IP address 1.1.1.1 during the last 1 hour.\"
	Reason *string `json:"reason,omitempty"`
	Threshold *RiskEvaluationDetailsUserVelocityByIpThreshold `json:"threshold,omitempty"`
	Velocity *RiskEvaluationDetailsUserVelocityByIpVelocity `json:"velocity,omitempty"`
}

// NewRiskEvaluationDetailsUserVelocityByIp instantiates a new RiskEvaluationDetailsUserVelocityByIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsUserVelocityByIp() *RiskEvaluationDetailsUserVelocityByIp {
	this := RiskEvaluationDetailsUserVelocityByIp{}
	return &this
}

// NewRiskEvaluationDetailsUserVelocityByIpWithDefaults instantiates a new RiskEvaluationDetailsUserVelocityByIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsUserVelocityByIpWithDefaults() *RiskEvaluationDetailsUserVelocityByIp {
	this := RiskEvaluationDetailsUserVelocityByIp{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetLevel(v string) {
	o.Level = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetReason(v string) {
	o.Reason = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetThreshold() RiskEvaluationDetailsUserVelocityByIpThreshold {
	if o == nil || o.Threshold == nil {
		var ret RiskEvaluationDetailsUserVelocityByIpThreshold
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetThresholdOk() (*RiskEvaluationDetailsUserVelocityByIpThreshold, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given RiskEvaluationDetailsUserVelocityByIpThreshold and assigns it to the Threshold field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetThreshold(v RiskEvaluationDetailsUserVelocityByIpThreshold) {
	o.Threshold = &v
}

// GetVelocity returns the Velocity field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetVelocity() RiskEvaluationDetailsUserVelocityByIpVelocity {
	if o == nil || o.Velocity == nil {
		var ret RiskEvaluationDetailsUserVelocityByIpVelocity
		return ret
	}
	return *o.Velocity
}

// GetVelocityOk returns a tuple with the Velocity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetVelocityOk() (*RiskEvaluationDetailsUserVelocityByIpVelocity, bool) {
	if o == nil || o.Velocity == nil {
		return nil, false
	}
	return o.Velocity, true
}

// HasVelocity returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasVelocity() bool {
	if o != nil && o.Velocity != nil {
		return true
	}

	return false
}

// SetVelocity gets a reference to the given RiskEvaluationDetailsUserVelocityByIpVelocity and assigns it to the Velocity field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetVelocity(v RiskEvaluationDetailsUserVelocityByIpVelocity) {
	o.Velocity = &v
}

func (o RiskEvaluationDetailsUserVelocityByIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Velocity != nil {
		toSerialize["velocity"] = o.Velocity
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationDetailsUserVelocityByIp struct {
	value *RiskEvaluationDetailsUserVelocityByIp
	isSet bool
}

func (v NullableRiskEvaluationDetailsUserVelocityByIp) Get() *RiskEvaluationDetailsUserVelocityByIp {
	return v.value
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIp) Set(val *RiskEvaluationDetailsUserVelocityByIp) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsUserVelocityByIp) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsUserVelocityByIp(val *RiskEvaluationDetailsUserVelocityByIp) *NullableRiskEvaluationDetailsUserVelocityByIp {
	return &NullableRiskEvaluationDetailsUserVelocityByIp{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsUserVelocityByIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


