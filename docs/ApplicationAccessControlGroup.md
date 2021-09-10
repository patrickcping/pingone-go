# ApplicationAccessControlGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A string that specifies the group type required to access the application. Options are ANY_GROUP (the actor must belong to at least one group listed in the accessControl.group.groups property) and ALL_GROUPS (the actor must belong to all groups listed in the accessControl.group.groups property). | [optional] 
**Groups** | Pointer to **[]string** | A set that specifies the group IDs for the groups the actor must belong to for access to the application. | [optional] 

## Methods

### NewApplicationAccessControlGroup

`func NewApplicationAccessControlGroup() *ApplicationAccessControlGroup`

NewApplicationAccessControlGroup instantiates a new ApplicationAccessControlGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAccessControlGroupWithDefaults

`func NewApplicationAccessControlGroupWithDefaults() *ApplicationAccessControlGroup`

NewApplicationAccessControlGroupWithDefaults instantiates a new ApplicationAccessControlGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationAccessControlGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationAccessControlGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationAccessControlGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationAccessControlGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroups

`func (o *ApplicationAccessControlGroup) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ApplicationAccessControlGroup) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ApplicationAccessControlGroup) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ApplicationAccessControlGroup) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


