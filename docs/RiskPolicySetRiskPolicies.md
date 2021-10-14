# RiskPolicySetRiskPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**RiskPolicySetCondition**](RiskPolicySetCondition.md) |  | 
**CreatedAt** | Pointer to **string** | The time the resource was first created (format ISO-8061). | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies a description for this risk policy. This is an optional property. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, punctuation character, or space. Maximum size is 1024 characters. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies a name for this risk policy. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. Maximum size is 256 characters. | 
**Priority** | Pointer to **int32** | An integer that specifies priority of the policy inside a risk policy set, designating which policy should run first. This is a read-only value. The priority is determined by the order in which policies are listed in the policy set. The first policy in the list is assigned priority 1 and is evaluated first. The next policy in the list is assigned priority 2 and so on. | [optional] 
**Result** | [**RiskPolicyResult**](RiskPolicyResult.md) |  | 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewRiskPolicySetRiskPolicies

`func NewRiskPolicySetRiskPolicies(condition RiskPolicySetCondition, name string, result RiskPolicyResult, ) *RiskPolicySetRiskPolicies`

NewRiskPolicySetRiskPolicies instantiates a new RiskPolicySetRiskPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetRiskPoliciesWithDefaults

`func NewRiskPolicySetRiskPoliciesWithDefaults() *RiskPolicySetRiskPolicies`

NewRiskPolicySetRiskPoliciesWithDefaults instantiates a new RiskPolicySetRiskPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RiskPolicySetRiskPolicies) GetCondition() RiskPolicySetCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPolicySetRiskPolicies) GetConditionOk() (*RiskPolicySetCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPolicySetRiskPolicies) SetCondition(v RiskPolicySetCondition)`

SetCondition sets Condition field to given value.


### GetCreatedAt

`func (o *RiskPolicySetRiskPolicies) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPolicySetRiskPolicies) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPolicySetRiskPolicies) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPolicySetRiskPolicies) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *RiskPolicySetRiskPolicies) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPolicySetRiskPolicies) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPolicySetRiskPolicies) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPolicySetRiskPolicies) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPolicySetRiskPolicies) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPolicySetRiskPolicies) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPolicySetRiskPolicies) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPolicySetRiskPolicies) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *RiskPolicySetRiskPolicies) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPolicySetRiskPolicies) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPolicySetRiskPolicies) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPolicySetRiskPolicies) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPolicySetRiskPolicies) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPolicySetRiskPolicies) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPolicySetRiskPolicies) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *RiskPolicySetRiskPolicies) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RiskPolicySetRiskPolicies) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RiskPolicySetRiskPolicies) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RiskPolicySetRiskPolicies) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetResult

`func (o *RiskPolicySetRiskPolicies) GetResult() RiskPolicyResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskPolicySetRiskPolicies) GetResultOk() (*RiskPolicyResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskPolicySetRiskPolicies) SetResult(v RiskPolicyResult)`

SetResult sets Result field to given value.


### GetUpdatedAt

`func (o *RiskPolicySetRiskPolicies) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPolicySetRiskPolicies) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPolicySetRiskPolicies) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPolicySetRiskPolicies) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


