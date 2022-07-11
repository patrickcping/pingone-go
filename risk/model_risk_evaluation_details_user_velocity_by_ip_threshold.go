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

// RiskEvaluationDetailsUserVelocityByIpThreshold The information about the calculated threshold used.
type RiskEvaluationDetailsUserVelocityByIpThreshold struct {
	// An integer indicating the value calculated for the high threshold. If the IP was accessed by more than the high number of users during the past hour, the IP is flagged as a HIGH userVelocityByIp.level.
	High *int32 `json:"high,omitempty"`
	// An integer indicating the value calculated for the medium threshold. If the IP was accessed by more than the medium number of users during the past hour, the IP is flagged as a MEDIUM userVelocityByIp.level
	Medium *int32 `json:"medium,omitempty"`
	// An enum indicating the source used to calculate the threshold. This can be: MIN_NOT_REACHED. If the measure is less than every.minSample, the threshold isn't calculated. Instead, a value of LOW is automatically assigned. CALCULATED. Indicates the threshold was calculated. ENVIRONMENT_FALLBACK. Indicates a global threshold calculated for the entire environment is used. The global threshold is used when the userVelocityByIp.threshold couldn't be calculated for the user, generally due to a lack of past transactions for the risk predictor to use for the threshold calculation. DEFAULT_FALLBACK. Indicates the default threshold defined for the predictor (in threshold.medium or threshold.high) is used. The default threshold is used when ENVIRONMENT_FALLBACK (the global threshold) couldn't be calculated, generally due to a lack of past transactions for the risk predictor to use for the global threshold calculation.
	Source *string `json:"source,omitempty"`
	// A date-time indicating the timestamp for the calculated threshold.
	CalculatedAt *string `json:"calculatedAt,omitempty"`
	// A date-time indicating when the threshold will be recalculated. The recalculation will happen before this time.
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewRiskEvaluationDetailsUserVelocityByIpThreshold instantiates a new RiskEvaluationDetailsUserVelocityByIpThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsUserVelocityByIpThreshold() *RiskEvaluationDetailsUserVelocityByIpThreshold {
	this := RiskEvaluationDetailsUserVelocityByIpThreshold{}
	return &this
}

// NewRiskEvaluationDetailsUserVelocityByIpThresholdWithDefaults instantiates a new RiskEvaluationDetailsUserVelocityByIpThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsUserVelocityByIpThresholdWithDefaults() *RiskEvaluationDetailsUserVelocityByIpThreshold {
	this := RiskEvaluationDetailsUserVelocityByIpThreshold{}
	return &this
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetHigh() int32 {
	if o == nil || o.High == nil {
		var ret int32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetHighOk() (*int32, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given int32 and assigns it to the High field.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetHigh(v int32) {
	o.High = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetMedium() int32 {
	if o == nil || o.Medium == nil {
		var ret int32
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetMediumOk() (*int32, bool) {
	if o == nil || o.Medium == nil {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasMedium() bool {
	if o != nil && o.Medium != nil {
		return true
	}

	return false
}

// SetMedium gets a reference to the given int32 and assigns it to the Medium field.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetMedium(v int32) {
	o.Medium = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetSource(v string) {
	o.Source = &v
}

// GetCalculatedAt returns the CalculatedAt field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetCalculatedAt() string {
	if o == nil || o.CalculatedAt == nil {
		var ret string
		return ret
	}
	return *o.CalculatedAt
}

// GetCalculatedAtOk returns a tuple with the CalculatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetCalculatedAtOk() (*string, bool) {
	if o == nil || o.CalculatedAt == nil {
		return nil, false
	}
	return o.CalculatedAt, true
}

// HasCalculatedAt returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasCalculatedAt() bool {
	if o != nil && o.CalculatedAt != nil {
		return true
	}

	return false
}

// SetCalculatedAt gets a reference to the given string and assigns it to the CalculatedAt field.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetCalculatedAt(v string) {
	o.CalculatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o RiskEvaluationDetailsUserVelocityByIpThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.Medium != nil {
		toSerialize["medium"] = o.Medium
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.CalculatedAt != nil {
		toSerialize["calculatedAt"] = o.CalculatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationDetailsUserVelocityByIpThreshold struct {
	value *RiskEvaluationDetailsUserVelocityByIpThreshold
	isSet bool
}

func (v NullableRiskEvaluationDetailsUserVelocityByIpThreshold) Get() *RiskEvaluationDetailsUserVelocityByIpThreshold {
	return v.value
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIpThreshold) Set(val *RiskEvaluationDetailsUserVelocityByIpThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsUserVelocityByIpThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIpThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsUserVelocityByIpThreshold(val *RiskEvaluationDetailsUserVelocityByIpThreshold) *NullableRiskEvaluationDetailsUserVelocityByIpThreshold {
	return &NullableRiskEvaluationDetailsUserVelocityByIpThreshold{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsUserVelocityByIpThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIpThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

