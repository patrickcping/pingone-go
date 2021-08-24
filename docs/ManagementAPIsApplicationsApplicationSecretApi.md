# \ManagementAPIsApplicationsApplicationSecretApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDApplicationsAppIDSecretGet**](ManagementAPIsApplicationsApplicationSecretApi.md#V1EnvironmentsEnvIDApplicationsAppIDSecretGet) | **Get** /v1/environments/{envID}/applications/{appID}/secret | READ Application Secret
[**V1EnvironmentsEnvIDApplicationsAppIDSecretPost**](ManagementAPIsApplicationsApplicationSecretApi.md#V1EnvironmentsEnvIDApplicationsAppIDSecretPost) | **Post** /v1/environments/{envID}/applications/{appID}/secret | UPDATE Application Secret



## V1EnvironmentsEnvIDApplicationsAppIDSecretGet

> V1EnvironmentsEnvIDApplicationsAppIDSecretGet(ctx, envID, appID).Authorization(authorization).Execute()

READ Application Secret



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
    appID := "appID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationSecretApi.V1EnvironmentsEnvIDApplicationsAppIDSecretGet(context.Background(), envID, appID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSecretApi.V1EnvironmentsEnvIDApplicationsAppIDSecretGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSecretGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDApplicationsAppIDSecretPost

> V1EnvironmentsEnvIDApplicationsAppIDSecretPost(ctx, envID, appID).ContentType(contentType).Authorization(authorization).Execute()

UPDATE Application Secret



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
    appID := "appID_example" // string | 
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsApplicationsApplicationSecretApi.V1EnvironmentsEnvIDApplicationsAppIDSecretPost(context.Background(), envID, appID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsApplicationsApplicationSecretApi.V1EnvironmentsEnvIDApplicationsAppIDSecretPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDApplicationsAppIDSecretPostRequest struct via the builder pattern


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

