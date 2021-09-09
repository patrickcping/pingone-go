# BillOfMaterials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Products** | Pointer to [**[]BillOfMaterialsProducts**](BillOfMaterialsProducts.md) | An array that specifies the products associated with this bill of materials | [optional] 

## Methods

### NewBillOfMaterials

`func NewBillOfMaterials() *BillOfMaterials`

NewBillOfMaterials instantiates a new BillOfMaterials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfMaterialsWithDefaults

`func NewBillOfMaterialsWithDefaults() *BillOfMaterials`

NewBillOfMaterialsWithDefaults instantiates a new BillOfMaterials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillOfMaterials) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillOfMaterials) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillOfMaterials) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillOfMaterials) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillOfMaterials) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillOfMaterials) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillOfMaterials) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillOfMaterials) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProducts

`func (o *BillOfMaterials) GetProducts() []BillOfMaterialsProducts`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *BillOfMaterials) GetProductsOk() (*[]BillOfMaterialsProducts, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *BillOfMaterials) SetProducts(v []BillOfMaterialsProducts)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *BillOfMaterials) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


