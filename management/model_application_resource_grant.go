/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ApplicationResourceGrant struct for ApplicationResourceGrant
type ApplicationResourceGrant struct {
	Application *ApplicationResourceGrantApplication `json:"application,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies the application resource grant ID.
	Id *string `json:"id,omitempty"`
	Resource ApplicationResourceGrantResource `json:"resource"`
	Scopes []ApplicationResourceGrantScopesInner `json:"scopes"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewApplicationResourceGrant instantiates a new ApplicationResourceGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourceGrant(resource ApplicationResourceGrantResource, scopes []ApplicationResourceGrantScopesInner) *ApplicationResourceGrant {
	this := ApplicationResourceGrant{}
	this.Resource = resource
	this.Scopes = scopes
	return &this
}

// NewApplicationResourceGrantWithDefaults instantiates a new ApplicationResourceGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceGrantWithDefaults() *ApplicationResourceGrant {
	this := ApplicationResourceGrant{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ApplicationResourceGrant) GetApplication() ApplicationResourceGrantApplication {
	if o == nil || o.Application == nil {
		var ret ApplicationResourceGrantApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrant) GetApplicationOk() (*ApplicationResourceGrantApplication, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ApplicationResourceGrant) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ApplicationResourceGrantApplication and assigns it to the Application field.
func (o *ApplicationResourceGrant) SetApplication(v ApplicationResourceGrantApplication) {
	o.Application = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationResourceGrant) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrant) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationResourceGrant) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApplicationResourceGrant) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationResourceGrant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrant) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationResourceGrant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationResourceGrant) SetId(v string) {
	o.Id = &v
}

// GetResource returns the Resource field value
func (o *ApplicationResourceGrant) GetResource() ApplicationResourceGrantResource {
	if o == nil {
		var ret ApplicationResourceGrantResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrant) GetResourceOk() (*ApplicationResourceGrantResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ApplicationResourceGrant) SetResource(v ApplicationResourceGrantResource) {
	o.Resource = v
}

// GetScopes returns the Scopes field value
func (o *ApplicationResourceGrant) GetScopes() []ApplicationResourceGrantScopesInner {
	if o == nil {
		var ret []ApplicationResourceGrantScopesInner
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrant) GetScopesOk() ([]ApplicationResourceGrantScopesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ApplicationResourceGrant) SetScopes(v []ApplicationResourceGrantScopesInner) {
	o.Scopes = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationResourceGrant) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrant) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationResourceGrant) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ApplicationResourceGrant) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ApplicationResourceGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if true {
		toSerialize["scopes"] = o.Scopes
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationResourceGrant struct {
	value *ApplicationResourceGrant
	isSet bool
}

func (v NullableApplicationResourceGrant) Get() *ApplicationResourceGrant {
	return v.value
}

func (v *NullableApplicationResourceGrant) Set(val *ApplicationResourceGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourceGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourceGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourceGrant(val *ApplicationResourceGrant) *NullableApplicationResourceGrant {
	return &NullableApplicationResourceGrant{value: val, isSet: true}
}

func (v NullableApplicationResourceGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourceGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

