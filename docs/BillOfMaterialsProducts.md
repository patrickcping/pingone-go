# BillOfMaterialsProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the BOM ID | [optional] [readonly] 
**Type** | Pointer to **string** | A string that specifies the Ping Identity product type. Options for PingOne platform products are PING_ONE_MFA, PING_ONE_RISK, PING_ONE_VERIFY, and PING_ONE_BASE. The PING_ONE_BASE product represents the default set of services that an environment can use on the PingOne platform. Options for other Ping Identity products are PING_FEDERATE, PING_ACCESS, PING_DIRECTORY, PING_DATA_SYNC, PING_DATA_GOVERNANCE, PING_ONE_FOR_ENTERPRISE, PING_ID, PING_ID_SDK, PING_INTELLIGENCE, and PING_CENTRAL | [optional] 
**Description** | Pointer to **string** | A string that specifies the description of the product or standalone service | [optional] [readonly] 
**Console** | Pointer to [**BillOfMaterialsConsole**](BillOfMaterialsConsole.md) |  | [optional] 
**Bookmarks** | Pointer to [**[]BillOfMaterialsBookmarks**](BillOfMaterialsBookmarks.md) | Optional array of custom bookmarks. Maximum of five bookmarks per product. | [optional] 

## Methods

### NewBillOfMaterialsProducts

`func NewBillOfMaterialsProducts() *BillOfMaterialsProducts`

NewBillOfMaterialsProducts instantiates a new BillOfMaterialsProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfMaterialsProductsWithDefaults

`func NewBillOfMaterialsProductsWithDefaults() *BillOfMaterialsProducts`

NewBillOfMaterialsProductsWithDefaults instantiates a new BillOfMaterialsProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillOfMaterialsProducts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfMaterialsProducts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfMaterialsProducts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfMaterialsProducts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BillOfMaterialsProducts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillOfMaterialsProducts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillOfMaterialsProducts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillOfMaterialsProducts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BillOfMaterialsProducts) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfMaterialsProducts) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfMaterialsProducts) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfMaterialsProducts) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConsole

`func (o *BillOfMaterialsProducts) GetConsole() BillOfMaterialsConsole`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *BillOfMaterialsProducts) GetConsoleOk() (*BillOfMaterialsConsole, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *BillOfMaterialsProducts) SetConsole(v BillOfMaterialsConsole)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *BillOfMaterialsProducts) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetBookmarks

`func (o *BillOfMaterialsProducts) GetBookmarks() []BillOfMaterialsBookmarks`

GetBookmarks returns the Bookmarks field if non-nil, zero value otherwise.

### GetBookmarksOk

`func (o *BillOfMaterialsProducts) GetBookmarksOk() (*[]BillOfMaterialsBookmarks, bool)`

GetBookmarksOk returns a tuple with the Bookmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarks

`func (o *BillOfMaterialsProducts) SetBookmarks(v []BillOfMaterialsBookmarks)`

SetBookmarks sets Bookmarks field to given value.

### HasBookmarks

`func (o *BillOfMaterialsProducts) HasBookmarks() bool`

HasBookmarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


