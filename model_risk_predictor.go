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

// RiskPredictor struct for RiskPredictor
type RiskPredictor struct {
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights.
	Name *string `json:"name,omitempty"`
	// A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details).
	CompactName *string `json:"compactName,omitempty"`
	// A string type. This specifies the desription of the risk predictor. Maximum length is 1024 characters.
	Type *string `json:"type,omitempty"`
	// A string type. This specifies the desription of the risk predictor. Maximum length is 1024 characters.
	Description *string `json:"description,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Default *RiskPredictorDefault `json:"default,omitempty"`
	Map *RiskPredictorMap `json:"map,omitempty"`
}

// NewRiskPredictor instantiates a new RiskPredictor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictor() *RiskPredictor {
	this := RiskPredictor{}
	return &this
}

// NewRiskPredictorWithDefaults instantiates a new RiskPredictor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorWithDefaults() *RiskPredictor {
	this := RiskPredictor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPredictor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPredictor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPredictor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RiskPredictor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RiskPredictor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RiskPredictor) SetName(v string) {
	o.Name = &v
}

// GetCompactName returns the CompactName field value if set, zero value otherwise.
func (o *RiskPredictor) GetCompactName() string {
	if o == nil || o.CompactName == nil {
		var ret string
		return ret
	}
	return *o.CompactName
}

// GetCompactNameOk returns a tuple with the CompactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetCompactNameOk() (*string, bool) {
	if o == nil || o.CompactName == nil {
		return nil, false
	}
	return o.CompactName, true
}

// HasCompactName returns a boolean if a field has been set.
func (o *RiskPredictor) HasCompactName() bool {
	if o != nil && o.CompactName != nil {
		return true
	}

	return false
}

// SetCompactName gets a reference to the given string and assigns it to the CompactName field.
func (o *RiskPredictor) SetCompactName(v string) {
	o.CompactName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPredictor) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPredictor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskPredictor) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPredictor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPredictor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPredictor) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPredictor) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPredictor) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskPredictor) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPredictor) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPredictor) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RiskPredictor) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPredictor) GetDefault() RiskPredictorDefault {
	if o == nil || o.Default == nil {
		var ret RiskPredictorDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetDefaultOk() (*RiskPredictorDefault, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPredictor) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given RiskPredictorDefault and assigns it to the Default field.
func (o *RiskPredictor) SetDefault(v RiskPredictorDefault) {
	o.Default = &v
}

// GetMap returns the Map field value if set, zero value otherwise.
func (o *RiskPredictor) GetMap() RiskPredictorMap {
	if o == nil || o.Map == nil {
		var ret RiskPredictorMap
		return ret
	}
	return *o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictor) GetMapOk() (*RiskPredictorMap, bool) {
	if o == nil || o.Map == nil {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *RiskPredictor) HasMap() bool {
	if o != nil && o.Map != nil {
		return true
	}

	return false
}

// SetMap gets a reference to the given RiskPredictorMap and assigns it to the Map field.
func (o *RiskPredictor) SetMap(v RiskPredictorMap) {
	o.Map = &v
}

func (o RiskPredictor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CompactName != nil {
		toSerialize["compactName"] = o.CompactName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Map != nil {
		toSerialize["map"] = o.Map
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPredictor struct {
	value *RiskPredictor
	isSet bool
}

func (v NullableRiskPredictor) Get() *RiskPredictor {
	return v.value
}

func (v *NullableRiskPredictor) Set(val *RiskPredictor) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictor) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictor(val *RiskPredictor) *NullableRiskPredictor {
	return &NullableRiskPredictor{value: val, isSet: true}
}

func (v NullableRiskPredictor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


