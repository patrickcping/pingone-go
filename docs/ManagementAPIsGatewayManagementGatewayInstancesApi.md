# \ManagementAPIsGatewayManagementGatewayInstancesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDGatewaysGatewayIDInstancesGet**](ManagementAPIsGatewayManagementGatewayInstancesApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDInstancesGet) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/instances | READ All Gateway Instances
[**V1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGet**](ManagementAPIsGatewayManagementGatewayInstancesApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGet) | **Get** /v1/environments/{envID}/gateways/{gatewayID}/instances/{instanceID} | READ One Gateway Instance



## V1EnvironmentsEnvIDGatewaysGatewayIDInstancesGet

> V1EnvironmentsEnvIDGatewaysGatewayIDInstancesGet(ctx, envID, gatewayID).Execute()

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
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayInstancesApi.V1EnvironmentsEnvIDGatewaysGatewayIDInstancesGet(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayInstancesApi.V1EnvironmentsEnvIDGatewaysGatewayIDInstancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGet

> V1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGet(ctx, envID, gatewayID, instanceID).Execute()

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
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayInstancesApi.V1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGet(context.Background(), envID, gatewayID, instanceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayInstancesApi.V1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDInstancesInstanceIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

