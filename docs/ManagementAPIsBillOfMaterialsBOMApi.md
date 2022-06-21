# \ManagementAPIsBillOfMaterialsBOMApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadOneBillOfMaterials**](ManagementAPIsBillOfMaterialsBOMApi.md#ReadOneBillOfMaterials) | **Get** /v1/environments/{envID}/billOfMaterials | READ One Bill of Materials
[**UpdateBillOfMaterials**](ManagementAPIsBillOfMaterialsBOMApi.md#UpdateBillOfMaterials) | **Put** /v1/environments/{envID}/billOfMaterials | UPDATE Bill of Materials



## ReadOneBillOfMaterials

> BillOfMaterials ReadOneBillOfMaterials(ctx, envID).Execute()

READ One Bill of Materials



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsBillOfMaterialsBOMApi.ReadOneBillOfMaterials(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBillOfMaterialsBOMApi.ReadOneBillOfMaterials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneBillOfMaterials`: BillOfMaterials
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsBillOfMaterialsBOMApi.ReadOneBillOfMaterials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillOfMaterials**](BillOfMaterials.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillOfMaterials

> BillOfMaterials UpdateBillOfMaterials(ctx, envID).BillOfMaterials(billOfMaterials).Execute()

UPDATE Bill of Materials



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    envID := "envID_example" // string | 
    billOfMaterials := *openapiclient.NewBillOfMaterials() // BillOfMaterials |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsBillOfMaterialsBOMApi.UpdateBillOfMaterials(context.Background(), envID).BillOfMaterials(billOfMaterials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsBillOfMaterialsBOMApi.UpdateBillOfMaterials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBillOfMaterials`: BillOfMaterials
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsBillOfMaterialsBOMApi.UpdateBillOfMaterials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billOfMaterials** | [**BillOfMaterials**](BillOfMaterials.md) |  | 

### Return type

[**BillOfMaterials**](BillOfMaterials.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

