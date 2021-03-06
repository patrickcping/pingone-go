# RiskPolicySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time the resource was created (format ISO-8061). | [optional] [readonly] 
**Default** | Pointer to **bool** | A boolean that specifies whether this risk policy set is the environment&#39;s default risk policy set, which is used whenever an explicit policySet ID is not specified in the risk policy evaluation request. If this property is not specified, the value defaults to false, and this risk policy set is not regarded as the default risk policy set for the environment. When this property is set to true (in PUT or POST requests), the default property of all other risk policy sets in the environment is set to false. | [optional] 
**DefaultResult** | Pointer to [**RiskPolicyResult**](RiskPolicyResult.md) |  | [optional] 
**Description** | Pointer to **string** | A string that specifies a description for this policy set. This is an optional property. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, punctuation character, or space. Maximum size is 1024 characters. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies a name for this policy set. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. Maximum size is 256 characters. | 
**RiskPolicies** | Pointer to [**[]RiskPolicySetRiskPoliciesInner**](RiskPolicySetRiskPoliciesInner.md) | An array of policies related to this policy set. | [optional] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewRiskPolicySet

`func NewRiskPolicySet(name string, ) *RiskPolicySet`

NewRiskPolicySet instantiates a new RiskPolicySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetWithDefaults

`func NewRiskPolicySetWithDefaults() *RiskPolicySet`

NewRiskPolicySetWithDefaults instantiates a new RiskPolicySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RiskPolicySet) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPolicySet) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPolicySet) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPolicySet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *RiskPolicySet) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RiskPolicySet) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RiskPolicySet) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RiskPolicySet) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultResult

`func (o *RiskPolicySet) GetDefaultResult() RiskPolicyResult`

GetDefaultResult returns the DefaultResult field if non-nil, zero value otherwise.

### GetDefaultResultOk

`func (o *RiskPolicySet) GetDefaultResultOk() (*RiskPolicyResult, bool)`

GetDefaultResultOk returns a tuple with the DefaultResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultResult

`func (o *RiskPolicySet) SetDefaultResult(v RiskPolicyResult)`

SetDefaultResult sets DefaultResult field to given value.

### HasDefaultResult

`func (o *RiskPolicySet) HasDefaultResult() bool`

HasDefaultResult returns a boolean if a field has been set.

### GetDescription

`func (o *RiskPolicySet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPolicySet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPolicySet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPolicySet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPolicySet) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPolicySet) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPolicySet) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPolicySet) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *RiskPolicySet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPolicySet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPolicySet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPolicySet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPolicySet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPolicySet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPolicySet) SetName(v string)`

SetName sets Name field to given value.


### GetRiskPolicies

`func (o *RiskPolicySet) GetRiskPolicies() []RiskPolicySetRiskPoliciesInner`

GetRiskPolicies returns the RiskPolicies field if non-nil, zero value otherwise.

### GetRiskPoliciesOk

`func (o *RiskPolicySet) GetRiskPoliciesOk() (*[]RiskPolicySetRiskPoliciesInner, bool)`

GetRiskPoliciesOk returns a tuple with the RiskPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskPolicies

`func (o *RiskPolicySet) SetRiskPolicies(v []RiskPolicySetRiskPoliciesInner)`

SetRiskPolicies sets RiskPolicies field to given value.

### HasRiskPolicies

`func (o *RiskPolicySet) HasRiskPolicies() bool`

HasRiskPolicies returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskPolicySet) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPolicySet) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPolicySet) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPolicySet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


