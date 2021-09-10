# UserPopulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the identifier of the population resource associated with the user. This property cannot be updated using PUT {{apiPath}}/environments/{{envID}}/users/{{userID}}. However, it can be updated using PUT /environments/{{envID}}/users/{{userID}}/population. | [optional] 

## Methods

### NewUserPopulation

`func NewUserPopulation() *UserPopulation`

NewUserPopulation instantiates a new UserPopulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPopulationWithDefaults

`func NewUserPopulationWithDefaults() *UserPopulation`

NewUserPopulationWithDefaults instantiates a new UserPopulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPopulation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPopulation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPopulation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserPopulation) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


