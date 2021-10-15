# \ManagementAPIsGatewayManagementGatewayInstancesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllGatewayInstances**](ManagementAPIsGatewayManagementGatewayInstancesApi.md#ReadAllGatewayInstances) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/instances | READ All Gateway Instances
[**ReadOneGatewayInstance**](ManagementAPIsGatewayManagementGatewayInstancesApi.md#ReadOneGatewayInstance) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/instances/{instanceID} | READ One Gateway Instance



## ReadAllGatewayInstances

> EntityArray ReadAllGatewayInstances(ctx, envID, gatewayID).Execute()

READ All Gateway Instances



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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayInstancesApi.ReadAllGatewayInstances(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayInstancesApi.ReadAllGatewayInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGatewayInstances`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayInstancesApi.ReadAllGatewayInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllGatewayInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneGatewayInstance

> GatewayInstance ReadOneGatewayInstance(ctx, envID, gatewayID, instanceID).Execute()

READ One Gateway Instance



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
    gatewayID := "gatewayID_example" // string | 
    instanceID := "instanceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayInstancesApi.ReadOneGatewayInstance(context.Background(), envID, gatewayID, instanceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayInstancesApi.ReadOneGatewayInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGatewayInstance`: GatewayInstance
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayInstancesApi.ReadOneGatewayInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 
**instanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GatewayInstance**](GatewayInstance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

