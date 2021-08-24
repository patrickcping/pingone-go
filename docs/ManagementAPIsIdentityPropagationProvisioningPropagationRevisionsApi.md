# \ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet**](ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.md#V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet) | **Get** /v1/environments/{envID}/propagation/revisions/id:latest | READ Latest Revision
[**V1EnvironmentsEnvIDPropagationRevisionsPost**](ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.md#V1EnvironmentsEnvIDPropagationRevisionsPost) | **Post** /v1/environments/{envID}/propagation/revisions | CREATE Revision
[**V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet**](ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.md#V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet) | **Get** /v1/environments/{envID}/propagation/revisions/{previousRevisionID} | READ Previous Revision



## V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet

> V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet(ctx, envID).Accept(accept).Authorization(authorization).Execute()

READ Latest Revision



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
    accept := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet(context.Background(), envID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.V1EnvironmentsEnvIDPropagationRevisionsIdlatestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRevisionsIdlatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
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


## V1EnvironmentsEnvIDPropagationRevisionsPost

> V1EnvironmentsEnvIDPropagationRevisionsPost(ctx, envID).ContentType(contentType).Authorization(authorization).Execute()

CREATE Revision



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
    contentType := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.V1EnvironmentsEnvIDPropagationRevisionsPost(context.Background(), envID).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.V1EnvironmentsEnvIDPropagationRevisionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRevisionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet

> V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet(ctx, envID, previousRevisionID).Accept(accept).Authorization(authorization).Execute()

READ Previous Revision



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
    previousRevisionID := "previousRevisionID_example" // string | 
    accept := "application/json" // string |  (optional)
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet(context.Background(), envID, previousRevisionID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPIsIdentityPropagationProvisioningPropagationRevisionsApi.V1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**previousRevisionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDPropagationRevisionsPreviousRevisionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
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

