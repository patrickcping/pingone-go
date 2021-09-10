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

// EntityArrayEmbedded struct for EntityArrayEmbedded
type EntityArrayEmbedded struct {
	Applications *[]AnyOfApplicationSAMLApplicationOIDC `json:"applications,omitempty"`
	Environments *[]Environment `json:"environments,omitempty"`
	Groups *[]Group `json:"groups,omitempty"`
	GroupMemberships *[]GroupMembership `json:"groupMemberships,omitempty"`
	Populations *[]Population `json:"populations,omitempty"`
	Resources *[]Resource `json:"resources,omitempty"`
	RoleAssignments *[]RoleAssignment `json:"roleAssignments,omitempty"`
	Roles *[]Role `json:"roles,omitempty"`
	Users *[]User `json:"users,omitempty"`
}

// NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityArrayEmbedded() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetApplications() []AnyOfApplicationSAMLApplicationOIDC {
	if o == nil || o.Applications == nil {
		var ret []AnyOfApplicationSAMLApplicationOIDC
		return ret
	}
	return *o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetApplicationsOk() (*[]AnyOfApplicationSAMLApplicationOIDC, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []AnyOfApplicationSAMLApplicationOIDC and assigns it to the Applications field.
func (o *EntityArrayEmbedded) SetApplications(v []AnyOfApplicationSAMLApplicationOIDC) {
	o.Applications = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetEnvironments() []Environment {
	if o == nil || o.Environments == nil {
		var ret []Environment
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetEnvironmentsOk() (*[]Environment, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *EntityArrayEmbedded) SetEnvironments(v []Environment) {
	o.Environments = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGroupsOk() (*[]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *EntityArrayEmbedded) SetGroups(v []Group) {
	o.Groups = &v
}

// GetGroupMemberships returns the GroupMemberships field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGroupMemberships() []GroupMembership {
	if o == nil || o.GroupMemberships == nil {
		var ret []GroupMembership
		return ret
	}
	return *o.GroupMemberships
}

// GetGroupMembershipsOk returns a tuple with the GroupMemberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGroupMembershipsOk() (*[]GroupMembership, bool) {
	if o == nil || o.GroupMemberships == nil {
		return nil, false
	}
	return o.GroupMemberships, true
}

// HasGroupMemberships returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGroupMemberships() bool {
	if o != nil && o.GroupMemberships != nil {
		return true
	}

	return false
}

// SetGroupMemberships gets a reference to the given []GroupMembership and assigns it to the GroupMemberships field.
func (o *EntityArrayEmbedded) SetGroupMemberships(v []GroupMembership) {
	o.GroupMemberships = &v
}

// GetPopulations returns the Populations field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetPopulations() []Population {
	if o == nil || o.Populations == nil {
		var ret []Population
		return ret
	}
	return *o.Populations
}

// GetPopulationsOk returns a tuple with the Populations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetPopulationsOk() (*[]Population, bool) {
	if o == nil || o.Populations == nil {
		return nil, false
	}
	return o.Populations, true
}

// HasPopulations returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasPopulations() bool {
	if o != nil && o.Populations != nil {
		return true
	}

	return false
}

// SetPopulations gets a reference to the given []Population and assigns it to the Populations field.
func (o *EntityArrayEmbedded) SetPopulations(v []Population) {
	o.Populations = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetResources() []Resource {
	if o == nil || o.Resources == nil {
		var ret []Resource
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetResourcesOk() (*[]Resource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Resource and assigns it to the Resources field.
func (o *EntityArrayEmbedded) SetResources(v []Resource) {
	o.Resources = &v
}

// GetRoleAssignments returns the RoleAssignments field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRoleAssignments() []RoleAssignment {
	if o == nil || o.RoleAssignments == nil {
		var ret []RoleAssignment
		return ret
	}
	return *o.RoleAssignments
}

// GetRoleAssignmentsOk returns a tuple with the RoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRoleAssignmentsOk() (*[]RoleAssignment, bool) {
	if o == nil || o.RoleAssignments == nil {
		return nil, false
	}
	return o.RoleAssignments, true
}

// HasRoleAssignments returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRoleAssignments() bool {
	if o != nil && o.RoleAssignments != nil {
		return true
	}

	return false
}

// SetRoleAssignments gets a reference to the given []RoleAssignment and assigns it to the RoleAssignments field.
func (o *EntityArrayEmbedded) SetRoleAssignments(v []RoleAssignment) {
	o.RoleAssignments = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRoles() []Role {
	if o == nil || o.Roles == nil {
		var ret []Role
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRolesOk() (*[]Role, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *EntityArrayEmbedded) SetRoles(v []Role) {
	o.Roles = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *EntityArrayEmbedded) SetUsers(v []User) {
	o.Users = &v
}

func (o EntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.GroupMemberships != nil {
		toSerialize["groupMemberships"] = o.GroupMemberships
	}
	if o.Populations != nil {
		toSerialize["populations"] = o.Populations
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.RoleAssignments != nil {
		toSerialize["roleAssignments"] = o.RoleAssignments
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableEntityArrayEmbedded struct {
	value *EntityArrayEmbedded
	isSet bool
}

func (v NullableEntityArrayEmbedded) Get() *EntityArrayEmbedded {
	return v.value
}

func (v *NullableEntityArrayEmbedded) Set(val *EntityArrayEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbedded(val *EntityArrayEmbedded) *NullableEntityArrayEmbedded {
	return &NullableEntityArrayEmbedded{value: val, isSet: true}
}

func (v NullableEntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


