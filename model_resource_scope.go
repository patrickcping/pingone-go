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

// ResourceScope struct for ResourceScope
type ResourceScope struct {
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the resource scope name.
	Name string `json:"name"`
	// A string that specifies the description of the scope.
	Description *string `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// An array that specifies the user schema attributes that can be read or updated for the specified PingOne access control scope. The value is an array of schema attribute paths (such as username, name.given, shirtSize) that the scope controls. This property is supported only for the p1:read:user, p1:update:user and p1:read:user:{suffix} and p1:update:user:{suffix} scopes. No other PingOne platform scopes allow this behavior. Any attributes not listed in the attribute array are excluded from the read or update action. The wildcard path (*) in the array includes all attributes and cannot be used in conjunction with any other user schema attribute paths
	SchemaAttributes *[]string `json:"schemaAttributes,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewResourceScope instantiates a new ResourceScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceScope(name string) *ResourceScope {
	this := ResourceScope{}
	this.Name = name
	return &this
}

// NewResourceScopeWithDefaults instantiates a new ResourceScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceScopeWithDefaults() *ResourceScope {
	this := ResourceScope{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceScope) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceScope) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceScope) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ResourceScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceScope) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceScope) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceScope) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceScope) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ResourceScope) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ResourceScope) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ResourceScope) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetSchemaAttributes returns the SchemaAttributes field value if set, zero value otherwise.
func (o *ResourceScope) GetSchemaAttributes() []string {
	if o == nil || o.SchemaAttributes == nil {
		var ret []string
		return ret
	}
	return *o.SchemaAttributes
}

// GetSchemaAttributesOk returns a tuple with the SchemaAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetSchemaAttributesOk() (*[]string, bool) {
	if o == nil || o.SchemaAttributes == nil {
		return nil, false
	}
	return o.SchemaAttributes, true
}

// HasSchemaAttributes returns a boolean if a field has been set.
func (o *ResourceScope) HasSchemaAttributes() bool {
	if o != nil && o.SchemaAttributes != nil {
		return true
	}

	return false
}

// SetSchemaAttributes gets a reference to the given []string and assigns it to the SchemaAttributes field.
func (o *ResourceScope) SetSchemaAttributes(v []string) {
	o.SchemaAttributes = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ResourceScope) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ResourceScope) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ResourceScope) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ResourceScope) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceScope) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ResourceScope) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ResourceScope) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ResourceScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.SchemaAttributes != nil {
		toSerialize["schemaAttributes"] = o.SchemaAttributes
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableResourceScope struct {
	value *ResourceScope
	isSet bool
}

func (v NullableResourceScope) Get() *ResourceScope {
	return v.value
}

func (v *NullableResourceScope) Set(val *ResourceScope) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceScope) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceScope(val *ResourceScope) *NullableResourceScope {
	return &NullableResourceScope{value: val, isSet: true}
}

func (v NullableResourceScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


