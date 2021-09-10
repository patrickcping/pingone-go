# ApplicationMobileIntegrityDetectionCacheDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | An integer that specifies the number of minutes or hours that specify the duration between successful integrity detection calls. Every attestation request entails a certain time tradeoff. You can choose to cache successful integrity detection calls for a predefined duration, between a minimum of 1 minute and a maximum of 48 hours. If mobile.integrityDetection.mode is ENABLED, the cache duration must be set. | [optional] 
**Units** | Pointer to **string** | A string that specifies the time units of the mobile.integrityDetection.cacheDuration.amount :MINUTES, HOURS | [optional] 

## Methods

### NewApplicationMobileIntegrityDetectionCacheDuration

`func NewApplicationMobileIntegrityDetectionCacheDuration() *ApplicationMobileIntegrityDetectionCacheDuration`

NewApplicationMobileIntegrityDetectionCacheDuration instantiates a new ApplicationMobileIntegrityDetectionCacheDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationMobileIntegrityDetectionCacheDurationWithDefaults

`func NewApplicationMobileIntegrityDetectionCacheDurationWithDefaults() *ApplicationMobileIntegrityDetectionCacheDuration`

NewApplicationMobileIntegrityDetectionCacheDurationWithDefaults instantiates a new ApplicationMobileIntegrityDetectionCacheDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUnits

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ApplicationMobileIntegrityDetectionCacheDuration) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


