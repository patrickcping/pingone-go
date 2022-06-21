# \ManagementAPIsGatewayManagementGatewaysApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](ManagementAPIsGatewayManagementGatewaysApi.md#CreateGateway) | **Post** /v1/environments/{envID}/gateways | CREATE Gateway
[**DeleteGateway**](ManagementAPIsGatewayManagementGatewaysApi.md#DeleteGateway) | **Delete** /v1/environments/{envID}/gateways/{gatewayID} | DELETE Gateway
[**ReadAllGateways**](ManagementAPIsGatewayManagementGatewaysApi.md#ReadAllGateways) | **Get** /v1/environments/{envID}/gateways | READ All Gateways
[**ReadOneGateway**](ManagementAPIsGatewayManagementGatewaysApi.md#ReadOneGateway) | **Get** /v1/environments/{envID}/gateways/{gatewayID} | READ One Gateway
[**UpdateGateway**](ManagementAPIsGatewayManagementGatewaysApi.md#UpdateGateway) | **Put** /v1/environments/{envID}/gateways/{gatewayID} | UPDATE Gateway



## CreateGateway

> CreateGateway201Response CreateGateway(ctx, envID).CreateGatewayRequest(createGatewayRequest).Execute()

CREATE Gateway



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
    createGatewayRequest := openapiclient.createGateway_request{Gateway: openapiclient.NewGateway("Name_example", "Type_example", false)} // CreateGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewaysApi.CreateGateway(context.Background(), envID).CreateGatewayRequest(createGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewaysApi.CreateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGateway`: CreateGateway201Response
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewaysApi.CreateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGatewayRequest** | [**CreateGatewayRequest**](CreateGatewayRequest.md) |  | 

### Return type

[**CreateGateway201Response**](CreateGateway201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGateway

> DeleteGateway(ctx, envID, gatewayID).Execute()

DELETE Gateway



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewaysApi.DeleteGateway(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewaysApi.DeleteGateway``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGatewayRequest struct via the builder pattern


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


## ReadAllGateways

> EntityArray ReadAllGateways(ctx, envID).Execute()

READ All Gateways



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
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewaysApi.ReadAllGateways(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewaysApi.ReadAllGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGateways`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewaysApi.ReadAllGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllGatewaysRequest struct via the builder pattern


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


## ReadOneGateway

> CreateGateway201Response ReadOneGateway(ctx, envID, gatewayID).Execute()

READ One Gateway



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewaysApi.ReadOneGateway(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewaysApi.ReadOneGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGateway`: CreateGateway201Response
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewaysApi.ReadOneGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateGateway201Response**](CreateGateway201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGateway

> CreateGateway201Response UpdateGateway(ctx, envID, gatewayID).CreateGatewayRequest(createGatewayRequest).Execute()

UPDATE Gateway



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
    createGatewayRequest := openapiclient.createGateway_request{Gateway: openapiclient.NewGateway("Name_example", "Type_example", false)} // CreateGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewaysApi.UpdateGateway(context.Background(), envID, gatewayID).CreateGatewayRequest(createGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewaysApi.UpdateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGateway`: CreateGateway201Response
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewaysApi.UpdateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createGatewayRequest** | [**CreateGatewayRequest**](CreateGatewayRequest.md) |  | 

### Return type

[**CreateGateway201Response**](CreateGateway201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

