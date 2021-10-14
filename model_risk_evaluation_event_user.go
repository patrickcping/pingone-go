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

// RiskEvaluationEventUser struct for RiskEvaluationEventUser
type RiskEvaluationEventUser struct {
	// A string that specifies the ID of the user associated with the event (maximum size 1024 characters). This is a required property.
	Id string `json:"id"`
	// A string that specifies the name of the user associated with the event (maximum size 1024 characters).
	Name *string `json:"name,omitempty"`
	// A string that specifies the type of user associated with the event. Options are EXTERNAL. This is a required property.
	Type string `json:"type"`
	// An array of group names.
	Groups *[]RiskEvaluationEventUserGroups `json:"groups,omitempty"`
}

// NewRiskEvaluationEventUser instantiates a new RiskEvaluationEventUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationEventUser(id string, type_ string) *RiskEvaluationEventUser {
	this := RiskEvaluationEventUser{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewRiskEvaluationEventUserWithDefaults instantiates a new RiskEvaluationEventUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationEventUserWithDefaults() *RiskEvaluationEventUser {
	this := RiskEvaluationEventUser{}
	return &this
}

// GetId returns the Id field value
func (o *RiskEvaluationEventUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventUser) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RiskEvaluationEventUser) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RiskEvaluationEventUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RiskEvaluationEventUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RiskEvaluationEventUser) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *RiskEvaluationEventUser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventUser) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskEvaluationEventUser) SetType(v string) {
	o.Type = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *RiskEvaluationEventUser) GetGroups() []RiskEvaluationEventUserGroups {
	if o == nil || o.Groups == nil {
		var ret []RiskEvaluationEventUserGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventUser) GetGroupsOk() (*[]RiskEvaluationEventUserGroups, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *RiskEvaluationEventUser) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []RiskEvaluationEventUserGroups and assigns it to the Groups field.
func (o *RiskEvaluationEventUser) SetGroups(v []RiskEvaluationEventUserGroups) {
	o.Groups = &v
}

func (o RiskEvaluationEventUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationEventUser struct {
	value *RiskEvaluationEventUser
	isSet bool
}

func (v NullableRiskEvaluationEventUser) Get() *RiskEvaluationEventUser {
	return v.value
}

func (v *NullableRiskEvaluationEventUser) Set(val *RiskEvaluationEventUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationEventUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationEventUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationEventUser(val *RiskEvaluationEventUser) *NullableRiskEvaluationEventUser {
	return &NullableRiskEvaluationEventUser{value: val, isSet: true}
}

func (v NullableRiskEvaluationEventUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationEventUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


