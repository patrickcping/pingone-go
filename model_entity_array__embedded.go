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
	Attributes       *[]interface{}              `json:"attributes,omitempty"`
	Applications     *[]interface{}              `json:"applications,omitempty"`
	Credentials      *[]GatewayCredential        `json:"credentials,omitempty"`
	Environments     *[]Environment              `json:"environments,omitempty"`
	GatewayInstances *[]GatewayInstance          `json:"gatewayInstances,omitempty"`
	Gateways         *[]interface{}              `json:"gateways,omitempty"`
	Grants           *[]ApplicationResourceGrant `json:"grants,omitempty"`
	Groups           *[]Group                    `json:"groups,omitempty"`
	GroupMemberships *[]GroupMembership          `json:"groupMemberships,omitempty"`
	Populations      *[]Population               `json:"populations,omitempty"`
	PushCredentials  *[]interface{}              `json:"pushCredentials,omitempty"`
	Resources        *[]Resource                 `json:"resources,omitempty"`
	Scopes           *[]ResourceScope            `json:"scopes,omitempty"`
	RiskPolicySets   *[]RiskPolicySet            `json:"riskPolicySets,omitempty"`
	RiskPredictors   *[]RiskPredictor            `json:"riskPredictors,omitempty"`
	RoleAssignments  *[]RoleAssignment           `json:"roleAssignments,omitempty"`
	Roles            *[]Role                     `json:"roles,omitempty"`
	Schemas          *[]Schema                   `json:"schemas,omitempty"`
	Users            *[]User                     `json:"users,omitempty"`
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

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAttributes() []interface{} {
	if o == nil || o.Attributes == nil {
		var ret []interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAttributesOk() (*[]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []OneOfApplicationAttributeMappingSchemaAttributeResourceAttribute and assigns it to the Attributes field.
func (o *EntityArrayEmbedded) SetAttributes(v []interface{}) {
	o.Attributes = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetApplications() []interface{} {
	if o == nil || o.Applications == nil {
		var ret []interface{}
		return ret
	}
	return *o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetApplicationsOk() (*[]interface{}, bool) {
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
func (o *EntityArrayEmbedded) SetApplications(v []interface{}) {
	o.Applications = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetCredentials() []GatewayCredential {
	if o == nil || o.Credentials == nil {
		var ret []GatewayCredential
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetCredentialsOk() (*[]GatewayCredential, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []GatewayCredential and assigns it to the Credentials field.
func (o *EntityArrayEmbedded) SetCredentials(v []GatewayCredential) {
	o.Credentials = &v
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

// GetGatewayInstances returns the GatewayInstances field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGatewayInstances() []GatewayInstance {
	if o == nil || o.GatewayInstances == nil {
		var ret []GatewayInstance
		return ret
	}
	return *o.GatewayInstances
}

// GetGatewayInstancesOk returns a tuple with the GatewayInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGatewayInstancesOk() (*[]GatewayInstance, bool) {
	if o == nil || o.GatewayInstances == nil {
		return nil, false
	}
	return o.GatewayInstances, true
}

// HasGatewayInstances returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGatewayInstances() bool {
	if o != nil && o.GatewayInstances != nil {
		return true
	}

	return false
}

// SetGatewayInstances gets a reference to the given []GatewayInstance and assigns it to the GatewayInstances field.
func (o *EntityArrayEmbedded) SetGatewayInstances(v []GatewayInstance) {
	o.GatewayInstances = &v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGateways() []interface{} {
	if o == nil || o.Gateways == nil {
		var ret []interface{}
		return ret
	}
	return *o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGatewaysOk() (*[]interface{}, bool) {
	if o == nil || o.Gateways == nil {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGateways() bool {
	if o != nil && o.Gateways != nil {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []AnyOfGatewayGatewayLDAP and assigns it to the Gateways field.
func (o *EntityArrayEmbedded) SetGateways(v []interface{}) {
	o.Gateways = &v
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGrants() []ApplicationResourceGrant {
	if o == nil || o.Grants == nil {
		var ret []ApplicationResourceGrant
		return ret
	}
	return *o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGrantsOk() (*[]ApplicationResourceGrant, bool) {
	if o == nil || o.Grants == nil {
		return nil, false
	}
	return o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGrants() bool {
	if o != nil && o.Grants != nil {
		return true
	}

	return false
}

// SetGrants gets a reference to the given []ApplicationResourceGrant and assigns it to the Grants field.
func (o *EntityArrayEmbedded) SetGrants(v []ApplicationResourceGrant) {
	o.Grants = &v
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

// GetPushCredentials returns the PushCredentials field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetPushCredentials() []interface{} {
	if o == nil || o.PushCredentials == nil {
		var ret []interface{}
		return ret
	}
	return *o.PushCredentials
}

// GetPushCredentialsOk returns a tuple with the PushCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetPushCredentialsOk() (*[]interface{}, bool) {
	if o == nil || o.PushCredentials == nil {
		return nil, false
	}
	return o.PushCredentials, true
}

// HasPushCredentials returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasPushCredentials() bool {
	if o != nil && o.PushCredentials != nil {
		return true
	}

	return false
}

// SetPushCredentials gets a reference to the given []AnyOfMFAPushCredentialAPNSMFAPushCredential and assigns it to the PushCredentials field.
func (o *EntityArrayEmbedded) SetPushCredentials(v []interface{}) {
	o.PushCredentials = &v
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

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetScopes() []ResourceScope {
	if o == nil || o.Scopes == nil {
		var ret []ResourceScope
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetScopesOk() (*[]ResourceScope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ResourceScope and assigns it to the Scopes field.
func (o *EntityArrayEmbedded) SetScopes(v []ResourceScope) {
	o.Scopes = &v
}

// GetRiskPolicySets returns the RiskPolicySets field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRiskPolicySets() []RiskPolicySet {
	if o == nil || o.RiskPolicySets == nil {
		var ret []RiskPolicySet
		return ret
	}
	return *o.RiskPolicySets
}

// GetRiskPolicySetsOk returns a tuple with the RiskPolicySets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRiskPolicySetsOk() (*[]RiskPolicySet, bool) {
	if o == nil || o.RiskPolicySets == nil {
		return nil, false
	}
	return o.RiskPolicySets, true
}

// HasRiskPolicySets returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRiskPolicySets() bool {
	if o != nil && o.RiskPolicySets != nil {
		return true
	}

	return false
}

// SetRiskPolicySets gets a reference to the given []RiskPolicySet and assigns it to the RiskPolicySets field.
func (o *EntityArrayEmbedded) SetRiskPolicySets(v []RiskPolicySet) {
	o.RiskPolicySets = &v
}

// GetRiskPredictors returns the RiskPredictors field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRiskPredictors() []RiskPredictor {
	if o == nil || o.RiskPredictors == nil {
		var ret []RiskPredictor
		return ret
	}
	return *o.RiskPredictors
}

// GetRiskPredictorsOk returns a tuple with the RiskPredictors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRiskPredictorsOk() (*[]RiskPredictor, bool) {
	if o == nil || o.RiskPredictors == nil {
		return nil, false
	}
	return o.RiskPredictors, true
}

// HasRiskPredictors returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRiskPredictors() bool {
	if o != nil && o.RiskPredictors != nil {
		return true
	}

	return false
}

// SetRiskPredictors gets a reference to the given []RiskPredictor and assigns it to the RiskPredictors field.
func (o *EntityArrayEmbedded) SetRiskPredictors(v []RiskPredictor) {
	o.RiskPredictors = &v
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

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetSchemas() []Schema {
	if o == nil || o.Schemas == nil {
		var ret []Schema
		return ret
	}
	return *o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetSchemasOk() (*[]Schema, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []Schema and assigns it to the Schemas field.
func (o *EntityArrayEmbedded) SetSchemas(v []Schema) {
	o.Schemas = &v
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
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.GatewayInstances != nil {
		toSerialize["gatewayInstances"] = o.GatewayInstances
	}
	if o.Gateways != nil {
		toSerialize["gateways"] = o.Gateways
	}
	if o.Grants != nil {
		toSerialize["grants"] = o.Grants
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
	if o.PushCredentials != nil {
		toSerialize["pushCredentials"] = o.PushCredentials
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.RiskPolicySets != nil {
		toSerialize["riskPolicySets"] = o.RiskPolicySets
	}
	if o.RiskPredictors != nil {
		toSerialize["riskPredictors"] = o.RiskPredictors
	}
	if o.RoleAssignments != nil {
		toSerialize["roleAssignments"] = o.RoleAssignments
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
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
