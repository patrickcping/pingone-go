# \ManagementAPIsGatewayManagementGatewayCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGatewayCredential**](ManagementAPIsGatewayManagementGatewayCredentialsApi.md#CreateGatewayCredential) | **Post** /v1/environments/{envID}/gateways/{gatewayID}/credentials | CREATE Gateway Credentials
[**DeleteGatewayCredential**](ManagementAPIsGatewayManagementGatewayCredentialsApi.md#DeleteGatewayCredential) | **Delete** /v1/environments/{envID}/gateways/{gatewayID}/credentials/{credentialID} | DELETE Gateway Credentials



## CreateGatewayCredential

> GatewayCredential CreateGatewayCredential(ctx, envID, gatewayID).Execute()

CREATE Gateway Credentials



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
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayCredentialsApi.CreateGatewayCredential(context.Background(), envID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayCredentialsApi.CreateGatewayCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGatewayCredential`: GatewayCredential
    fmt.Fprintf(os.Stdout, "Response from `ManagementAPIsGatewayManagementGatewayCredentialsApi.CreateGatewayCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayCredential**](GatewayCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGatewayCredential

> DeleteGatewayCredential(ctx, envID, gatewayID, credentialID).Execute()

DELETE Gateway Credentials



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
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagementAPIsGatewayManagementGatewayCredentialsApi.DeleteGatewayCredential(context.Background(), envID, gatewayID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayCredentialsApi.DeleteGatewayCredential``: %v\n", err)
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
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayCredentialRequest struct via the builder pattern


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

