# RiskPolicySetCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Equals** | Pointer to [**OneOfstringboolean**](oneOf&lt;string,boolean&gt;.md) |  | [optional] 
**AggregatedWeights** | Pointer to [**[]RiskPolicySetConditionAggregatedWeights**](RiskPolicySetConditionAggregatedWeights.md) |  | [optional] 
**Between** | Pointer to [**RiskPolicySetConditionBetween**](RiskPolicySetConditionBetween.md) |  | [optional] 

## Methods

### NewRiskPolicySetCondition

`func NewRiskPolicySetCondition() *RiskPolicySetCondition`

NewRiskPolicySetCondition instantiates a new RiskPolicySetCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetConditionWithDefaults

`func NewRiskPolicySetConditionWithDefaults() *RiskPolicySetCondition`

NewRiskPolicySetConditionWithDefaults instantiates a new RiskPolicySetCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RiskPolicySetCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskPolicySetCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskPolicySetCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskPolicySetCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEquals

`func (o *RiskPolicySetCondition) GetEquals() OneOfstringboolean`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *RiskPolicySetCondition) GetEqualsOk() (*OneOfstringboolean, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *RiskPolicySetCondition) SetEquals(v OneOfstringboolean)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *RiskPolicySetCondition) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetAggregatedWeights

`func (o *RiskPolicySetCondition) GetAggregatedWeights() []RiskPolicySetConditionAggregatedWeights`

GetAggregatedWeights returns the AggregatedWeights field if non-nil, zero value otherwise.

### GetAggregatedWeightsOk

`func (o *RiskPolicySetCondition) GetAggregatedWeightsOk() (*[]RiskPolicySetConditionAggregatedWeights, bool)`

GetAggregatedWeightsOk returns a tuple with the AggregatedWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedWeights

`func (o *RiskPolicySetCondition) SetAggregatedWeights(v []RiskPolicySetConditionAggregatedWeights)`

SetAggregatedWeights sets AggregatedWeights field to given value.

### HasAggregatedWeights

`func (o *RiskPolicySetCondition) HasAggregatedWeights() bool`

HasAggregatedWeights returns a boolean if a field has been set.

### GetBetween

`func (o *RiskPolicySetCondition) GetBetween() RiskPolicySetConditionBetween`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *RiskPolicySetCondition) GetBetweenOk() (*RiskPolicySetConditionBetween, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *RiskPolicySetCondition) SetBetween(v RiskPolicySetConditionBetween)`

SetBetween sets Between field to given value.

### HasBetween

`func (o *RiskPolicySetCondition) HasBetween() bool`

HasBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


