# ApplicationMobileIntegrityDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | A string that specifies whether device integrity detection takes place on mobile devices, for the application&#39;s enrollment and authentication events ENABLED, DISABLED | [optional] 
**CacheDuration** | Pointer to [**ApplicationMobileIntegrityDetectionCacheDuration**](ApplicationMobileIntegrityDetectionCacheDuration.md) |  | [optional] 

## Methods

### NewApplicationMobileIntegrityDetection

`func NewApplicationMobileIntegrityDetection() *ApplicationMobileIntegrityDetection`

NewApplicationMobileIntegrityDetection instantiates a new ApplicationMobileIntegrityDetection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationMobileIntegrityDetectionWithDefaults

`func NewApplicationMobileIntegrityDetectionWithDefaults() *ApplicationMobileIntegrityDetection`

NewApplicationMobileIntegrityDetectionWithDefaults instantiates a new ApplicationMobileIntegrityDetection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ApplicationMobileIntegrityDetection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApplicationMobileIntegrityDetection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApplicationMobileIntegrityDetection) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApplicationMobileIntegrityDetection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCacheDuration

`func (o *ApplicationMobileIntegrityDetection) GetCacheDuration() ApplicationMobileIntegrityDetectionCacheDuration`

GetCacheDuration returns the CacheDuration field if non-nil, zero value otherwise.

### GetCacheDurationOk

`func (o *ApplicationMobileIntegrityDetection) GetCacheDurationOk() (*ApplicationMobileIntegrityDetectionCacheDuration, bool)`

GetCacheDurationOk returns a tuple with the CacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDuration

`func (o *ApplicationMobileIntegrityDetection) SetCacheDuration(v ApplicationMobileIntegrityDetectionCacheDuration)`

SetCacheDuration sets CacheDuration field to given value.

### HasCacheDuration

`func (o *ApplicationMobileIntegrityDetection) HasCacheDuration() bool`

HasCacheDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


