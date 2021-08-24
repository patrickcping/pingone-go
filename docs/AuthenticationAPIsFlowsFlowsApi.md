# \AuthenticationAPIsFlowsFlowsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvIDFlowsFlowIDGet**](AuthenticationAPIsFlowsFlowsApi.md#EnvIDFlowsFlowIDGet) | **Get** /{envID}/flows/{flowID} | READ Flow
[**V1EnvIDFlowsFlowIDPost**](AuthenticationAPIsFlowsFlowsApi.md#V1EnvIDFlowsFlowIDPost) | **Post** /v1/{envID}/flows/{flowID} | Reset Flow



## EnvIDFlowsFlowIDGet

> EnvIDFlowsFlowIDGet(ctx, envID, flowID).Cookie(cookie).Execute()

READ Flow



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
    flowID := "flowID_example" // string | 
    cookie := "ST={{sessionToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsFlowsFlowsApi.EnvIDFlowsFlowIDGet(context.Background(), envID, flowID).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsFlowsFlowsApi.EnvIDFlowsFlowIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvIDFlowsFlowIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cookie** | **string** |  | 

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


## V1EnvIDFlowsFlowIDPost

> V1EnvIDFlowsFlowIDPost(ctx, envID, flowID).ContentType(contentType).Cookie(cookie).Execute()

Reset Flow



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
    flowID := "flowID_example" // string | 
    contentType := "application/vnd.pingidentity.session.reset+json" // string |  (optional)
    cookie := "ST={{sessionToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationAPIsFlowsFlowsApi.V1EnvIDFlowsFlowIDPost(context.Background(), envID, flowID).ContentType(contentType).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIsFlowsFlowsApi.V1EnvIDFlowsFlowIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvIDFlowsFlowIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
 **cookie** | **string** |  | 

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

