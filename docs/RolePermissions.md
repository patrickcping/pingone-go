# RolePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** | A string that specifies the resource for which the permission is applicable. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of what the permission enables for the role. | [optional] [readonly] 

## Methods

### NewRolePermissions

`func NewRolePermissions() *RolePermissions`

NewRolePermissions instantiates a new RolePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsWithDefaults

`func NewRolePermissionsWithDefaults() *RolePermissions`

NewRolePermissionsWithDefaults instantiates a new RolePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *RolePermissions) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *RolePermissions) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *RolePermissions) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *RolePermissions) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetDescription

`func (o *RolePermissions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePermissions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePermissions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolePermissions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


