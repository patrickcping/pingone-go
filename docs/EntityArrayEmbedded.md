# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]ApplicationAttributeMapping**](ApplicationAttributeMapping.md) |  | [optional] 
**Applications** | Pointer to [**[]AnyOfApplicationSAMLApplicationOIDC**](AnyOfApplicationSAMLApplicationOIDC.md) |  | [optional] 
**Environments** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 
**Grants** | Pointer to [**[]ApplicationResourceGrant**](ApplicationResourceGrant.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**GroupMemberships** | Pointer to [**[]GroupMembership**](GroupMembership.md) |  | [optional] 
**Populations** | Pointer to [**[]Population**](Population.md) |  | [optional] 
**Resources** | Pointer to [**[]Resource**](Resource.md) |  | [optional] 
**RoleAssignments** | Pointer to [**[]RoleAssignment**](RoleAssignment.md) |  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *EntityArrayEmbedded) GetAttributes() []ApplicationAttributeMapping`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntityArrayEmbedded) GetAttributesOk() (*[]ApplicationAttributeMapping, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntityArrayEmbedded) SetAttributes(v []ApplicationAttributeMapping)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EntityArrayEmbedded) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetApplications

`func (o *EntityArrayEmbedded) GetApplications() []AnyOfApplicationSAMLApplicationOIDC`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EntityArrayEmbedded) GetApplicationsOk() (*[]AnyOfApplicationSAMLApplicationOIDC, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EntityArrayEmbedded) SetApplications(v []AnyOfApplicationSAMLApplicationOIDC)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *EntityArrayEmbedded) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetEnvironments

`func (o *EntityArrayEmbedded) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *EntityArrayEmbedded) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *EntityArrayEmbedded) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *EntityArrayEmbedded) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetGrants

`func (o *EntityArrayEmbedded) GetGrants() []ApplicationResourceGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *EntityArrayEmbedded) GetGrantsOk() (*[]ApplicationResourceGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *EntityArrayEmbedded) SetGrants(v []ApplicationResourceGrant)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *EntityArrayEmbedded) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetGroups

`func (o *EntityArrayEmbedded) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EntityArrayEmbedded) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EntityArrayEmbedded) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EntityArrayEmbedded) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetGroupMemberships

`func (o *EntityArrayEmbedded) GetGroupMemberships() []GroupMembership`

GetGroupMemberships returns the GroupMemberships field if non-nil, zero value otherwise.

### GetGroupMembershipsOk

`func (o *EntityArrayEmbedded) GetGroupMembershipsOk() (*[]GroupMembership, bool)`

GetGroupMembershipsOk returns a tuple with the GroupMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberships

`func (o *EntityArrayEmbedded) SetGroupMemberships(v []GroupMembership)`

SetGroupMemberships sets GroupMemberships field to given value.

### HasGroupMemberships

`func (o *EntityArrayEmbedded) HasGroupMemberships() bool`

HasGroupMemberships returns a boolean if a field has been set.

### GetPopulations

`func (o *EntityArrayEmbedded) GetPopulations() []Population`

GetPopulations returns the Populations field if non-nil, zero value otherwise.

### GetPopulationsOk

`func (o *EntityArrayEmbedded) GetPopulationsOk() (*[]Population, bool)`

GetPopulationsOk returns a tuple with the Populations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulations

`func (o *EntityArrayEmbedded) SetPopulations(v []Population)`

SetPopulations sets Populations field to given value.

### HasPopulations

`func (o *EntityArrayEmbedded) HasPopulations() bool`

HasPopulations returns a boolean if a field has been set.

### GetResources

`func (o *EntityArrayEmbedded) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EntityArrayEmbedded) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EntityArrayEmbedded) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EntityArrayEmbedded) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *EntityArrayEmbedded) GetRoleAssignments() []RoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *EntityArrayEmbedded) GetRoleAssignmentsOk() (*[]RoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *EntityArrayEmbedded) SetRoleAssignments(v []RoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *EntityArrayEmbedded) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### GetRoles

`func (o *EntityArrayEmbedded) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *EntityArrayEmbedded) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *EntityArrayEmbedded) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *EntityArrayEmbedded) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUsers

`func (o *EntityArrayEmbedded) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EntityArrayEmbedded) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EntityArrayEmbedded) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *EntityArrayEmbedded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


