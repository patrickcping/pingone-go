# \ManagementAPIsGatewayManagementGatewayCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDelete**](ManagementAPIsGatewayManagementGatewayCredentialsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDelete) | **Delete** /v1/environments/{envID}/gateways/{gatewayID}/credentials/{credentialID} | DELETE Gateway Credentials
[**V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPost**](ManagementAPIsGatewayManagementGatewayCredentialsApi.md#V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPost) | **Post** /v1/environments/{envID}/gateways/{gatewayID}/credentials | CREATE Gateway Credentials



## V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDelete

> V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDelete(ctx, envID, gatewayID, credentialID).Authorization(authorization).Execute()

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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayCredentialsApi.V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDelete(context.Background(), envID, gatewayID, credentialID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayCredentialsApi.V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDCredentialsCredentialIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPost

> V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPost(ctx, envID, gatewayID).ContentType(contentType).Authorization(authorization).Execute()

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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsGatewayManagementGatewayCredentialsApi.V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPost(context.Background(), envID, gatewayID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsGatewayManagementGatewayCredentialsApi.V1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDGatewaysGatewayIDCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

